print_folder_recurse() {
    for i in "$1"/*;do
        if [ -d "$i" ];then
            echo "dir: $i"
            print_folder_recurse "$i"
        elif [ -f "$i" ]; then
            case $i in *.go ) mv -- "$i" "$(expr "$i" : '\(.*\)\.go').tmpl" ;; esac
        fi
    done
}

print_folder_recurse modern