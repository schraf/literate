include(FetchContent)

function(literate_init GIT_TAG)
    FetchContent_Declare(
        literate
        GIT_REPOSITORY https://github.com/schraf/literate.git
        GIT_TAG        ${GIT_TAG}
    )
    FetchContent_MakeAvailable(literate)
    
    set(LITERATE_BIN "${CMAKE_BINARY_DIR}/bin/literate${CMAKE_EXECUTABLE_SUFFIX}" CACHE INTERNAL "")
    
    add_custom_command(
        OUTPUT "${LITERATE_BIN}"
        COMMAND go build -o "${LITERATE_BIN}" ./cmd/...
        WORKING_DIRECTORY "${literate_SOURCE_DIR}"
        COMMENT "Building literate tool..."
        VERBATIM
    )
    
    # Create a target for the executable so other steps can depend on it
    add_custom_target(build_literate DEPENDS "${LITERATE_BIN}")
endfunction()

function(literate_project TARGET_NAME)
    set(options "")
    set(oneValueArgs OUTPUT_DIR)
    set(multiValueArgs INPUT_FILES OUTPUTS DEPENDS)

    cmake_parse_arguments(ARG "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})

    set(LITERATE_COMMAND "${LITERATE_BIN}")
    if(ARG_OUTPUT_DIR)
        list(APPEND LITERATE_COMMAND -output "${ARG_OUTPUT_DIR}")
    endif()
    list(APPEND LITERATE_COMMAND ${ARG_INPUT_FILES})

    add_custom_command(
        OUTPUT ${ARG_OUTPUTS}
        COMMAND ${LITERATE_COMMAND}
        DEPENDS "${LITERATE_BIN}" ${ARG_INPUT_FILES} ${ARG_DEPENDS}
        WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
        COMMENT "Generating code for ${TARGET_NAME}"
        VERBATIM
    )

    target_sources(${TARGET_NAME} PRIVATE ${ARG_OUTPUTS})
    add_dependencies(${TARGET_NAME} build_literate)
endfunction()

