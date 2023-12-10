install(
    TARGETS aoc-2023_exe
    RUNTIME COMPONENT aoc-2023_Runtime
)

if(PROJECT_IS_TOP_LEVEL)
  include(CPack)
endif()
