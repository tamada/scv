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
        --format | -f)
            COMPREPLY=($(compgen -W "default xml json" -- "${cur}"))
            return 0
            ;;
        --input-type | -t)
            COMPREPLY=($(compgen -W "string json byte_file term_file" -- "${cur}"))
            return 0
            ;;
    esac
    opts="-a --algorithm -f --format -t --input-type -h --help"
    if [[ "$cur" =~ ^\- ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    else
        compopt -o filenames
        COMPREPLY=($(compgen -d -- "$cur"))
    fi
}

complete -F __scv -o bashdefault -o default scv
