# Go-Port-Scanner
**small and super quick concurrent port scanner using the GO language**

This scanner is not stealthy at all. it generates a massive amount of traffic. 

The UDP port scanner doesn't work too well because it only checks if it can connect. 
The TCP scanner works without a flaw.

# USAGE
To run the command from the terminal, use the following syntax:
> go run main.go target_ip Starting_port Ending_port protocol

If the protocol is not included, it will default to TCP:

**Example:**
> go run main.go 192.168.0.1 25564 25566

By default, closed ports are hidden. If you want to see closed ports, You can uncomment the lines in the connect function
