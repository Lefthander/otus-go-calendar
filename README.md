# otus-go-calendar
Study a clean architecture aproach for Go


TODO:
Phase 1:

- [ ] Create a skeleton of project in accordance with clean architecture aproach
- [ ] Create a data model of calendar
- [ ] Create a corresponding methods for data models
     - [ ] Add event to the store
     - [ ] Delete events from the store
     - [ ] Update events in the store
     - [ ] List all events in the store ( by userid or specific time)
- [ ] Create corresponding errors for business logic (Sentinel Errors)
- [ ] Write Unit Tests to verify business logic, in particular Sentinel Errors


Phase 2:

- [ ] Handling comand line parameters ( parameter --config 'filename')
- [ ] Read the configuration file.
      - The configuration file should contain the following parameters:
        - http_listen_port, http_ip
        - path to the log file
        - log level
- [ ] Create loggers with specified logging-level. Logging level provided by configuration file.
      - Logging levels : error / warning / info / debug   
- [ ] Create and run the basic - 'hello world' http server. Webserver should handle only /hello url.
