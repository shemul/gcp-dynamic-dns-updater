# GCP CloudDNS Updater

## About <a name = "about"></a>

By default, IPs are empirical unless you reserve IP. 
That means whenever you restart a GCP compute engine 
the machine will get a new ip from the DHCP pool. So your 
depended on services needs be updated according to the 
new IP. 
To emerge this problem we can set up a CloudDNS
that will actually point to the machine current IP. But
still you need to update the CloudDNS record whenever the 
machine rebooted.  

So this solution will works as a container using 
the command `update-dns` 
(entrypoint /app), 
host network and restart policy `always`. 
It will then sync the IP address automatically 
whenever it changes.

## Getting Started <a name = "getting_started"></a>

Quick Jump

`docker run -v /tmp/secret:/secret --net=host -e GOOGLE_APPLICATION_CREDENTIALS=/secert/service-account.json -e DNS_NAMES=app1.demo-yourdnszone.com. -e GOOGLE_PROJECT=<google-project-id> -it shemul/gcp-dynamic-dns-updater:latest update-dns`

to obtain the service account

` IAM & admin > Service accounts -> create a service account -> DNS > DNS Administrator(role)`

### Run Building
prepare and download the deps using

`go mod download`

I've added a build script to build 
the binary that actually we needs

`bash build.sh -v 1.0.0`

It will produce a docker image. 