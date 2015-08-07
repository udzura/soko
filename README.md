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

## Configuration

* Exists `/etc/soko.toml`

```toml
[default]
# This is a uri that stands for backend(consul, aws, openstack)
uri = "consul://foo.example.consul:8500/"

[openstack]
# OpenStack specific config
username = "udzura"
password = "XXXXXXXXXXXXXX"
tenant_name = "soko-test"
auth_url = "https://keystone.example.com:1234/v2.0"
region = "RegionOne"

[aws]
# AWS specific config
access_key_id = "AKIXXXXXXXXXXXXX"
secret_access_key = "4Jr6DXXXXXXXXXXXXXXXXXXXXXX"
region = "ap-northeast-1"
```

## soko will work on

* Cloud servers(such as EC2, OpenStack... with file `/var/lib/cloud/data/instance-id` existing)
* Cousul cluster backended

### Under testing

* OpenStack compute API v2
* AWS EC2

## Yaruzo!!

* Google Compute Engine
* etcd....???
* Redis....??????
