```
#include <loguru.hpp>

…

// Optional, but useful to time-stamp the start of the log.
// Will also detect verbosity level on command line as -v.
loguru::init(argc, argv);

// Put every log message in "everything.log":
loguru::add_file("everything.log", loguru::Append, loguru::Verbosity_MAX);

// Only log INFO, WARNING, ERROR and FATAL to "latest_readable.log":
loguru::add_file("latest_readable.log", loguru::Truncate, loguru::Verbosity_INFO);

// Only show most relevant things on stderr:
loguru::g_stderr_verbosity = 1;

LOG_SCOPE_F(INFO, "Will indent all log messages within this scope.");
LOG_F(INFO, "I'm hungry for some %.3f!", 3.14159);
LOG_F(2, "Will only show if verbosity is 2 or higher");
VLOG_F(get_log_level(), "Use vlog for dynamic log level (integer in the range 0-9, inclusive)");
LOG_IF_F(ERROR, badness, "Will only show if badness happens");
auto fp = fopen(filename, "r");
CHECK_F(fp != nullptr, "Failed to open file '%s'", filename);
CHECK_GT_F(length, 0); // Will print the value of `length` on failure.
CHECK_EQ_F(a, b, "You can also supply a custom message, like to print something: %d", a + b);

// Each function also comes with a version prefixed with D for Debug:
DCHECK_F(expensive_check(x)); // Only checked #if !NDEBUG
DLOG_F(INFO, "Only written in debug-builds");

// Turn off writing to stderr:
loguru::g_stderr_verbosity = loguru::Verbosity_OFF;

// Turn off writing err/warn in red:
loguru::g_colorlogtostderr = false;

// Throw exceptions instead of aborting on CHECK fails:
loguru::set_fatal_handler([](const loguru::Message& message){
	throw std::runtime_error(std::string(message.prefix) + message.message);
});
```

```
#define LOGURU_WITH_STREAMS 1
#include <loguru.hpp>
...
LOG_S(INFO) << "Look at my custom object: " << a.cross(b);
CHECK_EQ_S(pi, 3.14) << "Maybe it is closer to " << M_PI;
```