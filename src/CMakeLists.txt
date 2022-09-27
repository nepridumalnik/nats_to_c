set(TARGET main)

include_directories(${GO_NATS_INCLUDE})

message("-- GO_NATS_LIBRARIES: ${GO_NATS_LIBRARIES}")

add_executable(${TARGET} main.cpp)
target_link_libraries(${TARGET} ${GO_NATS_LIBRARIES})
