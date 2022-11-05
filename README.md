Server for upload video

### 梳理分片上传过程

**step 1: 对整个文件的操作**   (/api/v1/upload/media)
前端：根据上传的文件，计算其md5与大小，传给后端三个参数: `md5,size,filename`。
后端：根据传过来的参数,创建一个目录(可以直接引用md5值)，文件大小留着有用，文件名字用于保存分片时的命名。返回`id(文件标识符)`
**可能存在的问题：媒体信息的重复上传，即正在进行的文件再次上传 和 已经存在的文件再次上传。**

**step2: 文件分片操作**  (/api/v1/upload/chunk)
前端：对文件进行分片操作，每个切片大小根据协议调整，传参数:` id(文件ID),chunk_id(切片ID),md5,size(切片大小),file(文件流)`
后端：对传输过来的切片进行校验，并保存在对应的目录下
**可能存在的问题：分片重传**

**step3: 文件合并操作**  (/api/v1/upload/merge)
前端: 发送一个文件合并的请求，传`id(文件ID)与count(分片总数)`
后端：对收到的切片进行合并，并进行md5与size校验，返回值。
**可能存在的问题：文件ID不存在 、重传或 计算的md5值比较错误。**

**tip: 注意并发时操作的原子性**

**参考文献**

https://segmentfault.com/a/1190000040982815 ffmpeg使用
https://www.ruanyifeng.com/blog/2020/01/ffmpeg.html 