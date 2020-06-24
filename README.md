# ipu
ip utility 

### install

```
./install.sh
```

### help 

```
ipu -h
  -h
  -ip string
    	IPv4 Addresses, for example: 192.0.0.0 or 2002::1234:abcd:ffff:c0a8:100.
  -l	list all possible IP adddresses within a given a CIDR block.
  -p string
    	ports which each host should try to ping, for example: 80,22,5432 or 75-80
  -pto int
    	ping port timeout window in seconds.  (default 3)
  -sn string
    	CIDR subnet (IPv4/IPv6), for example: 192.0.0.0/8 or 2002::1234:abcd:ffff:c0a8:101/122.
  -v	current version (v0.4.0)
```

### examples

```
ipu -sn 192.0.0.0/8 -ip 192.0.0.0,193.0.0.1

CIDR Subnet Details
  Subnet:  192.0.0.0/8
  First IP:  192.0.0.0
  Last IP:  192.255.255.255
  Total Host:  16777216

Range Results
192.0.0.0 in subnet 192.0.0.0/8
193.0.0.1 not in subnet 192.0.0.0/8
```