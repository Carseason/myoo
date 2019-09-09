# Myoo - 好想告诉你
    后端使用 go语言 gin框架开发， orm使用beego的orm
    前端使用Vue开发。
    ps: 目前移动端自适应还没完成，当前版本只适合于桌面端

# 安装配置
    目前程序支持前后分离、服务端渲染单页模式。

    1.下载 myoo.zip 并解压, 赋予 myoo 二进制文件执行的权限 。

    2.配置 conf/config.json文件,
        Port 为程序端口,格式应为 :8080
        ServerRender 为服务端单页渲染,如果你需要前后分离把值改为 false 即可
        数据库默认使用mysql,请在 mysql 下创建 要使用的表 如 myoo 
            Db.Uname 数据库账号
            DB.Pass  数据库密码
            Db.Table 数据库表名称 
            Db.Addr  数据库地址,如 127.0.0.1:3306
        默认使用 Redis 作为缓存， 可不使用
            Redis.Addr  Redis地址,如 127.0.0.1:6379
            Redis.Pass Redis链接密码
        Key ---重要，作用于 jwt 加密 以及 程序对数据加密的密钥,请使用 16 位的密钥
    
    3.服务器渲染可跳过本步骤, 如使用前后分离的模式 ，需要在 views/index.html 的第15行 meta标签的content 其值填入你的后端服务器地址，
        如 <meta name="server" id="myoo-server" content="https://api.A.com">    域名结尾不要 /

    4.程序启动，程序使用二进制文件，在当前目录下执行 ./myoo 即可。

    5.使用 nginx 代理端口。nginx需要添加如下规则:
        location / {
            try_files /_not_exists_ @backend;
        }
        location @backend {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;
            proxy_pass http://127.0.0.1:8080;
        }

    6.前后分离需要配置nginx代理 views 目录以及 src  目录，把 views/index.html文件作为展示域名,src/ 为静态资源目录
        服务端 nginx 需要去除跨域限制，添加如下规则 ( * 可改为自己的域名 )
                location / {
                try_files /_not_exists_ @backend;
                add_header Access-Control-Allow-Origin *;
                add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
                add_header Access-Control-Allow-Headers 'X-MYOO';
                if ($request_method = 'OPTIONS') {
                    return 204;
                }
            }
        域名 A.com 使用 views/index.html ， src/* 目录作为前台
        域名 api.A.com 代理程序的8080端口 作为 api 服务端

    7.程序会在你数据库下创建如下表:
            category    //分类
            comments    //评论
            favorite    //文章收藏
            followers   //关注粉丝
            options     //配置信息
            posts       //文章
            posts_meta  //文章资源,如视频 音乐 下载信息
            posts_supported //文章点赞
            user        //用户

    8.程序启动完毕后,进入到注册页面 /sign/registered 注册用户，
        注册成功后需要把数据库 user 表的注册用户的 admin 字段改为 1 ,意为管理员，
        如果你配置了缓存信息 ，则需要 进入到 redis-cli 里 执行 del Myoo/User/1 命令删除当前用户缓存。
    
    9.配置网站信息, 进入 /admin 后台进行配置以及添加用户组,配置完成后可点击 删除缓存。

# 演示
   <a href="https://myoo.moelq.com" target="_blank">演示网站</a>

    
