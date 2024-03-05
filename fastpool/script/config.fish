if status is-interactive
    # Commands to run in interactive sessions can go here

    set fish_greeting # Disable greeting
    alias dc="docker compose"
    alias dcu="dc up -d"
    alias dcd="dc down"
    alias dcl="dc logs -f"
    alias dcp="dc pull"
    alias dcr="dc restart"
    alias dc-update="dcp && dcu"

    function ta
        tar -cvf $argv[1].tar $argv[1]
    end

    function targz
        tar -zcvf $argv[1].tar.gz $argv[1]
    end

    function untar
        tar -xvf $argv[1]
    end

    function untargz
        tar -zxvf $argv[1]
    end

    set PATH $HOME/.local/bin $HOME/.rye/shims /usr/local/go/bin /root/go/bin $PATH

end
