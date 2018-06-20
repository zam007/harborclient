# harborclient

用go编写，通过harbor API控制doker的工具；
harbor相关信息在conf/conf.go目录中填写；

## 清理harbor
#### 1.遍历所有repo，只保留最近30次tag
```
harborclient delete tag
```
#### 2.配合以下脚本，定时清理harbor仓库
```
#!/bin/bash

cd /home/worker/src/harbor

/bin/docker-compose stop

if [[ $? -eq 0 ]]
then
  /bin/docker run -it --name gc --rm --volumes-from registry vmware/registry:2.6.2-photon garbage-collect  /etc/registry/config.yml
  /bin/docker run --name gc --rm --volumes-from registry vmware/registry:2.6.2-photon garbage-collect  /etc/registry/config.yml
else
  echo "stop harbor faild, exit"
  exit 1
fi

/bin/docker-compose start 
```
#### 3.设置计划任务
```
# clear docker reg every week
59 04 * * 6 cd /root/scripts && /bin/bash -x clear_docker_reg.sh > clear_docker_reg.log 2>&1
# clear harbor repo tag 
59 3 * * * cd /root/scripts && nohup /root/scripts/harborclient delete tag > harborclient.log 2>&1 &
```

## 配合Jenkins获取repo版本号

```
harborclient get tag --repoName xxx
```
#### groovy 脚本
```
import groovy.json.JsonSlurper
def recurse 
def versionArraySort = { a1, a2 -> 
    def headCompare = a1[0] <=> a2[0] 
    if (a1.size() == 1 || a2.size() == 1 || headCompare != 0) { 
        return headCompare 
    } else { 
        return recurse(a1[1..-1], a2[1..-1]) 
    } 
} 
// fool Groovy to understand recursive closure 
recurse = versionArraySort
def versionStringSort = { s1, s2 -> 
    def nums = { it.tokenize('.').collect{ it.toInteger() } } 
    versionArraySort(nums(s1), nums(s2)) 
}
try {
    List<String> artifacts = new ArrayList<String>()
    def cmdapp = "/root/scripts/jenkinstag"
    def repo = "reserve/closer-h5"
    def artifactsObjectRaw = ["${cmdapp}","get","tag","--repoName","${repo}"].execute().text
    def jsonSlurper = new JsonSlurper()
    def artifactsJsonObject = jsonSlurper.parseText(artifactsObjectRaw)
    def dataArray = artifactsJsonObject.data
    for(item in dataArray){
        if (item.IsMetadata == false)
        artifacts.add(item.Text)
    } 
    return artifacts.sort(versionStringSort).reverse()
} catch (Exception e) {
    print "There was a problem fetching the artifacts" + e
}
```

