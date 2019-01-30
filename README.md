# system-stats-consul-template-plugin
Get system stats to use in Consul Template

Plugin based on [gopsutil](https://github.com/shirou/gopsutil) use system stats in Consul Template.
f.e. calculate HEAP for Java APPs. 

## Syntax

```
{{ plugin system-stats-consul-template-plugin "<statsType>"" }}
```

## Stats
 * memory
 * cpuInfo
 * cpuCount

## Expample
###get total mempory

```
{{ with $d := plugin "system-stats-consul-template-plugin" "memory" |parseJSON }}{{ $d.total }}{{ end }}
```

### Calculate HEAP Space in MB

```
{{ with $d := plugin "system-stats-consul-template-plugin" "memory" |parseJSON }}{{ printf "%.f" (($d.total | multiply 0.8) | divide 1000 | divide 1000) }}{{ end }}
```

## Links
 * [Consul Template](https://github.com/hashicorp/consul-template)
 * [gopsutil](https://github.com/shirou/gopsutil)
 
