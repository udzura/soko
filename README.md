# soko

Metadata Manager, backended with Consul, EC2, OpenStack Nova and so on

## Install

* Download linux binary zipball via https://github.com/udzura/soko/releases/latest
* Then extract

```bash
[cloud-user@www901 ~]$ ./soko help
Usage of ./soko:
  -server-id="f6809f77-e7b2-4d66-832e-XXXX": Target server's ID to get/put/delete. Defaults to cloud-init's server ID
[cloud-user@www901 ~]$ ./soko get Hi
Value for Hi seems to be empty.

[cloud-user@www901 ~]$ ./soko get Test
Value for Test seems to be empty.

[cloud-user@www901 ~]$ ./soko put Test Hello
OK
[cloud-user@www901 ~]$ ./soko get Test
Hello
[cloud-user@www901 ~]$ ./soko delete Test
OK
[cloud-user@www901 ~]$ ./soko get Test
Value for Test seems to be empty.
```

## soko will work on

* Cloud servers(such as EC2, OpenStack... with file `/var/lib/cloud/data/instance-id` existing)
* Cousul cluster backended

## Yaruzo!!

* OpenSTack metadata backend
* EC2 tags backend
* Redis....??????
