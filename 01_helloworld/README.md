## Prerequisites
Installing snap on Pop!_OS
Snap can be installed on Pop!_OS from the command line. Open Terminal from the Applications launcher and type the following:


sudo apt update 
sudo apt install snapd 
Either log out and back in again, or restart your system, to ensure snap’s paths are updated correctly.

To test your system, install the hello-world snap and make sure it runs correctly:

sudo snap install hello-world
hello-world 6.3 from Canonical✓ installed
hello-world
Hello World!

Install Go
sudo snap install go --classic
go version

![Diagrams](/../Go/images/go_cli.png) 
![Diagrams](/../Go/images/types_of_packages.png)
![Diagrams](/../Go/images/go_package.png)
