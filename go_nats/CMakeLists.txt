find_program(GO go REQUIRED)
message("-- GO: ${GO}")

set(GO_FLAGS -buildmode=c-archive)
set(GO_NAME nats_to_c)

if(WIN32)
    set(GOLIBNAME ${GO_NAME}.lib)
else()
    set(GOLIBNAME ${GO_NAME}.a)
endif()

# 
set(GO_NATS_LIBRARIES ${CMAKE_BINARY_DIR}/${GOLIBNAME} CACHE INTERNAL find_go_nats)
set(GO_NATS_INCLUDE ${CMAKE_BINARY_DIR}/include CACHE INTERNAL find_go_nats)

add_custom_target(${GO_NAME}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMAND ${GO} build -o ${GOLIBNAME} ${GO_FLAGS}
    COMMAND ${CMAKE_COMMAND} -E copy_if_different ${CMAKE_CURRENT_SOURCE_DIR}/${GOLIBNAME} ${CMAKE_BINARY_DIR}
    COMMAND ${CMAKE_COMMAND} -E copy_if_different ${CMAKE_CURRENT_SOURCE_DIR}/${GO_NAME}.h ${GO_NATS_INCLUDE}/${GO_NAME}.h
)
