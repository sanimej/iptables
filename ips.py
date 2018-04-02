import sys
import subprocess

command = 'iptables -w -A INPUT -p tcp -s %s -j ACCEPT'
count = 0
for i in range(1,255):
    for j in range(1,255):
        for k in range(1,255):
            for l in range(1,255):
                ip = str(i)+'.'+str(j)+'.'+str(k)+'.'+str(l)
                iptables = command % (ip)+'\n'
                #print(iptables)
                #f.write(iptables)
                subprocess.run(iptables, shell=True)
                count += 1
                if count == int(sys.argv[1]):
                    sys.exit()
