## iptables impact on packet path

server.go implements a bare minimal tcp server which accepts, closes the connection right away and reports the number of 
established connection in a 30 seconds window.

client.go connects to the server's port from three go routines

ips.py is to add a given number of iptables rules to the INPUT chain.

### server's connection acceptable rate from three runes and with 15000 iptables rules in the INPUT chain

without iptables rules - ~207075
with iptables rules    - ~52700 

client & server were in separate EC2 t2.medium instances, running Ubuntu 16.04
