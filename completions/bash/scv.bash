__scv() {
    local i cur prev opts cmds
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    cmd=""
    opts=""

    case "${prev}" in
        --algorithm | -a)
            COMPREPLY=($(compgen -W "simpson jaccard dice cosine pearson euclidean manhattan chebyshev levenshtein" -- "${cur}"))
            return 0
            ;;
        --input-type | -t)
            COMPREPLY=($(compgen -W "byte_file term_file string json" -- "${cur}"))
            return 0
            ;;
        --format | -f)
            COMPREPLY=($(compgen -W "default json xml" -- "${cur}"))
            return 0
            ;;
    esac
    opts=" -a -f -t -h  --algorithm --format --input-type --help"
    if [[ "$cur" =~ ^\- ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    else
        compopt -o filenames
        COMPREPLY=($(compgen -d -- "$cur"))
    fi
}

complete -F __scv -o bashdefault -o default scv
