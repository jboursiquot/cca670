```
$ ssh -i cca670.pem ec2-user@the-ip
The authenticity of host 'the-ip (the-ip)' can't be established.
ECDSA key fingerprint is SHA256:/XQgamJtgPSJs92CYERh9I+6P72cj85GQvMJ24SEr5w.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'the-ip' (ECDSA) to the list of known hosts.

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
No packages needed for security; 26 packages available
Run "sudo yum update" to apply all updates.
[ec2-user@ip-10-103-8-59 ~]$ token='eyJ0eXAiOiJKV...' installerName=cloudinsights-linux-au-installer-1.205.0 ci_au_version=1.205.0 && curl $proxy_auth_scheme -H "Authorization: Bearer $token" -o $installerName.tar.gz https://60e5b2a2-d959-45e5-8d31-0d870bd0e384.c02.cloudinsights.netapp.com/rest/v1/au/installers/linux/1.205.0 && tar -xzvf $installerName.tar.gz -C . && sudo /bin/bash -c "TOKEN=$token hproxy='$https_proxy' hproxy_auth_scheme=$proxy_auth_scheme ci_au_version=$ci_au_version ./cloudinsights-linux-au-installer-1.205.0/cloudinsights-install.sh"

...
                     Welcome to CloudInsights (R) ..
                      Acquisition Unit

                        ____    ___
                       / ___|  |___|
                      | |       | |
                      | |___    | |
                       \____|  |___|
  NetApp (R)
  Installation: /opt/netapp/cloudinsights
  Logs:         /opt/netapp/cloudinsights/logs -> /var/log/netapp/cloudinsights

  To control the CloudInsights service:
    sudo cloudinsights-service.sh --help
  To uninstall:
    sudo cloudinsights-uninstall.sh --help

1/8 Acquisition Unit Starting
2/8 Connecting to Cloud Insights
3/8 Sending Certificate-Signing Request
4/8 Logging in to Cloud Insights
5/8 Updating Security Settings
6/8 Downloading Data Collection Modules
7/8 Registering to Cloud Insights
8/8 Acquisition Unit Ready

Acquisition Unit has been installed successfully.
```

```
[ec2-user@ip-the-ip ~]$ uptime
 04:24:25 up  2:39,  1 user,  load average: 0.03, 0.09, 0.08
[ec2-user@ip-the-ip ~]$ stress --cpu 2 -v --timeout 30s
stress: info: [10817] dispatching hogs: 2 cpu, 0 io, 0 vm, 0 hdd
stress: dbug: [10817] using backoff sleep of 6000us
stress: dbug: [10817] setting timeout to 30s
stress: dbug: [10817] --> hogcpu worker 2 [10818] forked
stress: dbug: [10817] using backoff sleep of 3000us
stress: dbug: [10817] setting timeout to 30s
stress: dbug: [10817] --> hogcpu worker 1 [10819] forked
stress: dbug: [10817] <-- worker 10818 signalled normally
stress: dbug: [10817] <-- worker 10819 signalled normally
stress: info: [10817] successful run completed in 30s
[ec2-user@ip-the-ip ~]$ uptime
 04:25:34 up  2:40,  1 user,  load average: 0.80, 0.26, 0.14
[ec2-user@ip-the-ip ~]$
```