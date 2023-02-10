#!/bin/sh

USAGE() {
    echo "Please specify argument."
}

if [ $# = 0 ]; then
    USAGE
    exit 1
fi

DEVELOPMENT_ENV_FILE="./container.dev.env"

load_env_development() {
    echo "[development::load_env_development] from ${DEVELOPMENT_ENV_FILE}"
    set -a
    # shellcheck source=./container.dev.env
    . "${DEVELOPMENT_ENV_FILE}"
    set +a
    echo "[development::load_env_development] development port: ${DEVELOPMENT_PORT}"

}

development() {
    echo "[development]"
    load_env_development
    commit_hash=$(git log -1 --format="%H")
    ymds=$(date +%Y%m%d%s)
    image_name="vms-backend-development-${commit_hash}-${ymds}"
    echo "[development] building container ${image_name}"
    docker build -f Dockerfile.dev -t "${image_name}" .

    if [ "$(docker images -q ${image_name} 2>/dev/null)" = "" ]; then
        echo "FAIL: [development] Image ${image_name} not found. Please check Docker app."
        exit 1
    else
        echo "[development] Image ${image_name} found"
        echo "[development] Running image ${image_name}"
        docker run -it -v "$(pwd):/app" -e SERVER_INGRESS_PORT="${DEVELOPMENT_PORT}" --rm -p "${DEVELOPMENT_PORT}":"${DEVELOPMENT_PORT}"/tcp "${image_name}" /bin/sh
    fi
    exit 0
}

case "${1}" in
dev)
    development
    ;;
esac
