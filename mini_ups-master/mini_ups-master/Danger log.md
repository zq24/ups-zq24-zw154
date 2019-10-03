### _Back-end danger log_

4.17 -- (fixed) learn how to use gorm, a Go package which provides API for ORM implementation to database. Try commonly used operations such as insert, delete, lookup, create, update, foreign key

4.18 -- (fixed) fail to receive corresponding Protobuf message, and find the reason that network transmission requires adding the object size at head of the byte stream

4.19 -- (fixed) connect to Amazon server to simulate a round trip, and encounter some unexpected exceptions. Find the reason that we will pass some fields not defined in the protobuf message to the database utility function will cause panic (similar to abort in C) in Go

4.20 -- (fixed) when we start to deploy the servers in docker, we need to figure out some configuration and connection such as ports mapping and exposure issues

4.21 -- (fixed) sometimes, the world server or UPS server may exit in an abnormal; the explanation for this behavior is due to that the database may not be ready to accept connection, when servers are trying to connect to the database. After being refused to connect, the servers will exit directly.

In purpose of reducing the possibility of this misbehavior, we add a sleep time before servers start up.  

4.23 -- (fixed) request transmission scheme for not receiving ack back from the world. We try to add interface (the only way to implement polymorphism in Go) in .pb.go package directly to create a go routine checking if a request specified the sequence number is still in the database. Try and check whether or not there will be a problem. Meanwhile, we post this issue in google protobuf github and stackoverflow. we are firmly convinced this this operation is legal.

4.24 -- (fixed) final testing on combining both front-end and back-end in docker, fix some connection problems.

Front-end danger log

### _Front-end danger log_

4.19 -- (fixed) the issue we encountered was passing value through URL and created the table in HTML dynamically. Since we avoided serving files other than HTML, it was difficult to write all the logic in HTML.

4.23 & 4.24 -- (not fixed) It was tricky to serve static files such as CSS, images or JavaScript files using Golang. It was very straightforward at the beginning, we just used StripPrefix and FileSever func to handle our static files. However, this would result a serious problem: by accessing to root URL, our directory structure would be exposed to public as well, which we wanted to prevent this situation from happening. We searched and tried some solutions like creating our own custom http file server and reject which request accessing to directory path. However, the results we got were not ideal. Therefore, all the files we served in this project are just HTML templates.

