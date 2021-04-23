

### 代理

export HTTP_PROXY=http://127.0.0.1:58591; export HTTPS_PROXY=http://127.0.0.1:58591; export ALL_PROXY=socks5://127.0.0.1:51837


export HTTP_PROXY=http://172.17.224.1:7890; export HTTPS_PROXY=http://172.17.224.1:7890; export ALL_PROXY=socks5://127.0.0.1:1080

export NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,172.17.20.34



function proxy_on() {
    export NO_PROXY="localhost,127.0.0.1"
    export HTTP_PROXY="http://172.17.224.1:7890"
    export HTTPS_PROXY=$http_proxy
    echo "Proxy environment variable set."
}
function proxy_off(){
    unset HTTP_PROXY
    unset HTTPS_PROXY
    echo -e "Proxy environment variable removed."
}