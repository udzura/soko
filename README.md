# soko (倉庫)

Server metadata inventory Manager, backended with Consul, EC2, OpenStack Nova and so on

## Why soko?

In some case we need to know the server's role(web server? api server?),
or master/slave (in case of MySQL, Solr, &c.), or some kind of metadata.

`soko` is designed for handling such metadata from Server inventory in one liner.

`soko` is very friendly with shell scripts(which are easy and available in `cloud-init`),
`ohai` for Chef, `facter` for Puppet, `run_command` in Itamae, and so on.

Inventory backends with:

* Consul KV
* EC2's Tags
* Metadata attribute in OpenStack Compute API v2

And more, soon...

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
# backend should be in [consul, aws, openstack]
backend = "consul"

[consul]
# Set just a consul backend URL
url = "consul://foo.example.consul:8500/"

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

* You can configure soko by one liner:

```bash
$ soko open aws access_key_id=$AWS_ACCESS_KEY_ID secret_access_key=$AWS_SECRET_ACCESS_KEY region=ap-northeast-1
# It goes such a way with other backends
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
