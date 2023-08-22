#
# Generated by @zondax/cli
#
VERSION 0.7
FROM alpine:3.17

source-root:
    RUN mkdir -p /source-root
    COPY . /source-root
    SAVE ARTIFACT /source-root

GET_SOURCE_CODE:
    COMMAND
    COPY +source-root/source-root .

COLLECT_SCRIPTS_RECURSIVE:
    COMMAND
    ARG --required BUNDLE_VARIANT

    RUN mkdir -p /zondax

    # First copy global defaults
    COPY +source-root/source-root/_default /zondax

    # Then copy relative default and the bundle variant
    COPY --if-exists ./_default /zondax
    IF [ ! -z "$BUNDLE_VARIANT" ]
        COPY --if-exists ./${BUNDLE_VARIANT} /zondax
    END

    # Fix permissions in case something is not set correctly
    RUN chmod +x /zondax/*.sh
    RUN chmod +x /zondax/entrypoint.d/*
    RUN chmod +x /zondax/k8s/*
    RUN chmod +x /zondax/utils.d/*

    ENTRYPOINT ["/zondax/entrypoint.sh"]
    CMD ["",""]

# This shared command can be used to publish images in a standardized format
# this will publish images named as zondax/${CONTAINER_FULLNAME}-{FLEXTAGS}
PUBLISH_WITH_FLEXTAGS:
    COMMAND
    ARG --required CONTAINER_FULLNAME
    ARG --required EARTHLY_GIT_SHORT_HASH
    ARG EARTHLY_GIT_COMMIT_TIMESTAMP
    ARG EARTHLY_GIT_TAG

    ARG --required EARTHLY_GIT_BRANCH
    ARG GIT_BRANCH=$(echo ${EARTHLY_GIT_BRANCH////-})

    WAIT
        # This will detect there is already a tag (:) and will use a dash - instead
        ENV DELIMITER=':'
        IF echo "$CONTAINER_FULLNAME" | grep -q ":"
            ENV DELIMITER='-'
        END

        # Store images
        # Tag an image using git's tag        
        IF [ ! -z "$EARTHLY_GIT_TAG" ]
            SAVE IMAGE --push zondax/${CONTAINER_FULLNAME}${DELIMITER}${EARTHLY_GIT_TAG}
        END

        # Tag the image using the commit timestamp
        IF [ ! -z "$EARTHLY_GIT_COMMIT_TIMESTAMP" ]
            ARG TIMESTAMP=$(date -d @${EARTHLY_GIT_COMMIT_TIMESTAMP} +"%Y%m%d%H%M%S")
            SAVE IMAGE --push zondax/${CONTAINER_FULLNAME}${DELIMITER}T${TIMESTAMP}
        END

        # Tag an image using the short commit hash
        SAVE IMAGE --push zondax/${CONTAINER_FULLNAME}${DELIMITER}${EARTHLY_GIT_SHORT_HASH}

        # Tag an image using the current git branch
        SAVE IMAGE --push zondax/${CONTAINER_FULLNAME}${DELIMITER}${GIT_BRANCH}

        # This image will never get pushed but can be used for local testing
        # example: docker run --rm -it --entrypoint /bin/bash zondax/node-fil-mainnet:v1.19.0-latest
        SAVE IMAGE zondax/${CONTAINER_FULLNAME}${DELIMITER}latest
    END

# Build all targets with name earthly* directories
all:
    WORKDIR /introspection
    COPY . .
    FOR TARGET IN $(ls -1d ./earthly*)
        BUILD "./${TARGET}"+all
    END
