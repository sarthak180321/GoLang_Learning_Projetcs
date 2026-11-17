# Process Managment and Creation
    fork/exec --> ForkExec/Exec
    Kill-->sends signal to processes to pause resume and shutdown
    Wait4--> wait for child Process to finish and checks its Exit code
    Getpid -->get process ID of current Runnning Process
    Getppid --> get the process ID of the parent

# File system --> Everything is a File Descriptor
    Open/close --> Low-Level acess  to Open File devices or Socket
    Dup/Dup2 --> clones the File descriptor to creare Multiple sockets files and channels
