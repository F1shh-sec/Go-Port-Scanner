# Go-Port-Scanner
small and super quick concurrent port scanner using the GO language
Not stealthy at all. 
The udp port scanner doesnt work too well because it only checks if it can connect. The TCP works well

# USAGE
To run the command from the terminal, use the following syntax
> go run main.go protocol target_ip Starting_port Ending_port

Example:
> go run main.go tcp 192.168.0.1 1 2000