# btrdb-ls
This utility lists collections and streams in BTrDB. 

## Build

To start, get dependencies and build the binary with `go get && go build`.

## Use

Then create a YAML config file `myconfig.yml` with the server location:
``` yaml
server: my.btrdb.server:4410
prefix: pingthingsio # You can filter results with the prefix key
```
and run `btrdb-ls myconfig.yml`. 

You'll see output that looks like 
```
Collection name                      Stream count
pingthingsio/sensor1                 19
pingthingsio/sensor2                 16
pingthingsio/sensor3                 19
```


If only one collection is found, a detail view of the collection
and its streams with be output. We can accomplish this by making
our prefix more specific:

``` yaml
server: my.btrdb.server:4410
prefix: pingthingsio/sensor1
```
The new output will look like:
```
Collection: pingthingsio/sensor1:
Streams:
 * UUID: b66fa23a-4abd-53b7-99f9-651f5f2fa3b1
 * Tags:
     - name: R1HNG
 * Annontations:
     - None

 * UUID: bf29bd31-0d8f-56c6-af3c-33251adb1009
 * Tags:
     - name: F2MAG
 * Annontations:
     - None

```
