DR. TSUN

---------------

Dr. Tsun checks if a given mailbox exists.

USAGE

    go run main.go example.com sales "John doe"

Given the previous command line arguments, Dr. Tsun will check the
mailboxes sales@example.com and john.doe@example.com.

Due to the nature of some mail providers, Dr. Tsun may provide
false negatives (e.g., sales@example.com exists but Dr. Tsun cannot
confirm).
