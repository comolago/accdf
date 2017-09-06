# Automated Testsuite Framework
**Author**: Marco Antonio Carcano

**Date**: 31/08/2017

**Document Version**: 0.0.1

**Document Status**: Draft
## Overview
Automated Test Suite (**ATF**) is a framework that allows to **run testsuites or single tests on targets that can even scale to whole infrastructures**.

This first release is **focused onto \*nix systems** (any kind of Unix, AIX, Linux, Mac OS-x and so on), but _hopefully in next release it will be able to run tests also on Microsoft based platforms_.

In order to understand the rest of this document you should be familiar with the following terms:

* **Deliverable**: a service that you want to deliver.
* **Delivered**: a **Deliverable** that has partially or completely delivered.
* **TestProject**: a project document used _to test a **Delivered**_ (or **part of it**). To perform tests TestProjects instantiates **TestSuites** and **TestCases** dynamically.
* **Benchmark**: a **script** (**or** a **binary program**) along with an **XML document** (_Benchmark XML_) that **describes it and the parameters it can accept**, the **platforms it can run onto** as well the **required command interpreter** (_bash, java and so on_).
In practice this XML is the descriptor (quite similar to a _WSDL_ for a web service), so _different version of the benchmark have different Benchmark XML_. (of course it is wise to update the version only for when adding or removing parameter to benchmarks, not for bugfixes and security enhancements. Along with the XML document, a **DTD** and an **XSL** are provided - the latter **enable the rendering of the document as a human readable document** that can be read when someone wants to create a new TestCase that uses this benchmark, or simply to better understand what the benchmark does. As a side note, the XML contains also **one or more checksums** that are verified by the framework before performing the tests - this means that if a benchmark does not matches its checksum (for example because it has been tampered) is not copied neither executed on the target and a security alert is raised.
* **BenchmarksLibrary**: a **library of Benchmarks**. It is served by a restfull service that enable to perform Benchmark XML lookups and get back the XML describing it.
* **TestCase**: a test you want to **carry out onto a part of a Delivered**. It is the **smallest unit you can run**, and is **made of several TestSteps**. A TestCase is described by an **XML document** (_TestCase XML_) that contains a **reference to a particular Benchmark** and a **list of TestSteps**. Along with the XML document a **DTD** and an **XSL** are provided - the latter enable the **rendering of the document as a human readable document** that can be read when someone wants to understand what the TestCase actually does or needs information on how to interpreter TestCase results.
* **TestStep**: is the unit that **describes the how to run a benchmark instance**, supplying the right **parameters** and **values**. It is possible to feed the same running benchmark with **multiple test steps**: this means that benchmark **executes test steps one by one**. For each test steps benchmark **returns a value** (**0=OK**, **1=FAILED**, **2=WARNING** and if necessary an **message**). The results report includes the result of every TestStep, along with information like a **timestamp** when it run, the **hostname** where it run, the **project version** along with its **GIT commit ID** and so on.
 Here are a few examples of TestCase/Benchmark/TestStep:

		CheckLDAPServers (TestCase) - executed Benchmark is "checkconn.pl"
		   ldap01.domain.local 389 (TestStep)
		   ldap01.domain.local 636 (TestStep)
		   ldap02.domain.local 389 (TestStep)
		   ldap02.domain.local 636 (TestStep)
		   
		CheckRedHatSatellite (TestCase) - executed Benchmark is "checkconn.pl"
          satellite.domain.local 80  (testsep)
          satellite.domain.local 443  (testsep)
          
		Backup (TestCase) - executed Benchmark is "checkbackup.sh"
		   include /foo/bar  (testsep)
		   exclude /home  (testsep)

 The first 2 TestCases could be part of "Networking Environment" TestSuite.
 
 TestSteps parameters **can be both static or dynamic**: if dynamic parameters are used, the **matching values are picked-up using a lookup facility that searches into CMDB, topology file and a custom config file**.
* **TestSuite** a collection of _TestCases XML_. TestSuites are used as a way to create a **group of TestCases** that an be **listed into TestProjects** avoiding to list every single TestCase.
* **TestCasesLibrary** a **library of TestSuites**. It is served by a restfull service that enable to perform TestSuites or TestCase XML lookups and get back the XML describing them.

##Requisites
**ATF** is desigend using distribuited services architecure paradigm, and it leverages on Docker Containers.

### Mandatory
* **Docker** environment - _Kubernetes strongly recommended, but not mandatory_
* **SQL Server** - the only available choice for now is **PostgreSQL**
* **NoSQL server** - the only available choiche for now is **Elasticsearch**
* **Versioned Document Repository** - the only available choiche for now is **GIT**
* **HTTP based Artifact Repository** - any **web server is suitable**, for example Apache HTTP
* **Task Scheduler** - the only available choiche for now is **Quartz**
* **CMDB** - the only available choiche for now is **HP Maximo CMBD**
* **Archimate compatible topology files**
* **Memcached** - used to store ephemeral data such as authentication tokens

###Optional
**Schedulers**

* **Jenkins**

**Additional Storage Repositories**

* **HP Quality Center (ALM)**


## Components
### CLI
It is a command line usable to interact to the service. It requires the user to authenticate using username and password.
CLI can perform a lot of commands by connecting to the API router.

1. CLI sends an **authentication request** to the **API Router**, providing **username** and **password** - the outcome is the **authentication token**
2. CLI sends the **actual command it wants to be executed** along with the **username** and the **authentication token**
3. Step 2 is **repeated as many time it is required** to complete the action requested by the user.

_CLI is used by **Jenkins** to launch tests_.

### API Router
Provides the only available endpoint used by clients to use the service.

Messages can be of two kind:

* **Authentication Request**: message contains **username** and **password**, and is forwarded to Authenticator. The reply of Authenticator, that if authentication succeeded contains an **authentication token**, is then sent back to the client, otherwise a **500 error** is returned

* **Service Request**: this kind of messages is actually a request to use one of the provided services.

 Requestor must be already authenticated either by:

 * supplying **username** and an **authentication token**. They are both sent to Authenticator to verify their validity and if it **doesn't succeed a 500 error** is returned to the client.
 * supplying **only the username** - this works only **if external authentication has been set**: in this case API Router trusts that the client has already verified the username passed - so authentication is actually skipped. _Note that **this option requires TLS client authentication**_.
 
 Service requests **without** one of the preceeding **authentication** tokens are **rejected with a 500 error**.
 
### Authenticator
Provides **authentication and authorization services**.
#### Authentication
It is performed using an **external backend** (for example LDAP, Active Directory, and so on).

Authentication requests are of 2 types:

* **New authentication**: it must contain a **username** and a **password**, that are then checked against the authentication backend. If authentication succeeds, an **authentication token is returned**. **Authentication token is stored** temporary onto **Memcached** along with the **IP of the requestor**, the **username** and an **expire date**.
* **Extend existing authentication**: it must contain the **username** and the **authentication token**. If an authentication token for the same username and the same IP exists and has not expired yet, authentication succeeds and **expire date is extended**.

#### Authorization
It is performed using an **external backend** (for example LDAP, Active Directory, and so on).

An authorization request contains the **username** to be checked along with **the action it wants to perform**

When serving authorization requests

1. it lookups the configured external backend to **get the list of the groups the user belongs to**.
2. it looks up a **role-to-group map** to guess **what roles are matching with the groups the user belongs to**.
3. it **checks if the action requested by the client is actually granted to the role** (by looking up an **action-to-role** map).

#### Roles listing
A last kind of request this object can serve is to **list the roles the user is granted to**. This _could be used by thirdy party authorization providers to perform additional authorization_ - for example _Document Level RBAC_.

When serving role listing requests:

1. it lookups the configured external backend to **get the list of the groups** the user belongs to.
2. it looks up an internal **role-to-group map** to guess what roles are matching with the groups the user belongs to.

The resulting **role list is then returned to the requestor**.

### Library
This restfull object serves both as **TestCasesLibrary** as well as **BenchmarkLibrary**. It can be used by both **Job Creator** to gather data to build the job file as well as by a **web browser** to render XML information in human readable format (_API Router routes to it requests to **/Benchmarks** and **/TestCases URLs**_).

**Before** serving the request it **connects to Authenticator** and passess it the **username** and **authentication token** to ensure that the request is allowed - _**Job creator** has it own credentials)_. If it succeed, it **returns the requested document** or **list of documents**, otherwise it **fails with a 500 error**.

It can be used to fetch:

* XML of a **single Benchmark**
* an XML of the **whole Benchmark Library**
* XML of a **single TestCase**
* XML of a **whole TestSuite**
* an XML of the **whole TestCases Library**

**Fetched documents** are **processed** by a **stack of hooks** (_if configured_) before sending them to the caller: Library **sends the whole document data to the hooked service** and **get back the document data processed by the service**. **These data** - that optionally could have been modified by the hooked service, are then **sent to the next hooked service until the last one is reached**. **Eventually document data are returned to the client** (or addedd to the list of documents that is eventualy sent to the requestor (for example when requesting TestSuites). **If any** of the preceeding steps **returns an empty document** the fetch is considered void (so optionally **other hooked service are skipped**) and the **empty data is returned to the caller** or,  if the request is for a list of documents (like a TestSuite), execution pass to the next requested document.
DocumentHooks are **configured as a stack** in the service **configuration file**, and are **sequentially processed**.
_By using this kind of hook based design it is possible to create additional services, such as the example **Document Level RBAC**_.
### Topology Generator
This service **generates the topology XML** file fetching informations of the resources listed in **Project XML** file - _that contains a version attribute and is fetched from a GIT repository_.

Data sources are:

* **CMDB**
* **Archimate compatible topology**
* **Custom settings file**

Service **has to be run on demand** whenever something relevant to the TestProject is added/modified/removed in any of the preceeding sources. **Topology XML** files are eventually stored **versioned in the GIT** repository. Every time **Topology XML** file is updated it is compared with the previous version: if they differs version number is incremented in the new version, and it is committed. **Project XML** file is then updated **increasing the version of Topology XML** to use, and the **commit ID** of Topology XML is added to **Project XML**, that is committed too.

Service requests **should be authorized by Authenticator**


### Job Creator
Service requests **should be authorized by Authenticator**

This service **generates the Job** that will be _later used by **Job Scheduler**_ to schedule Jobs.

A **Job** is a **GIT versioned filesystem subtree** structured as follows:

* Root directory contains _job.xml_ (**Job XML**) file, that contains the informations like **what to execute and where** to execute it, along with the **Project** it belongs to and its **GIT commit ID**. 
* A **directory for each host** for which there are testcases to run
 * Each "host" directory contains a **directory for each TestCase instance**
     * Each "Test" directory has the following structure
     
				etc - contains benchmark configuration file
				bin - contains the downloader and execute command
				var - contains the report created by the benchmark
    
      Downloader is an helper **that is able to download the Benchmark from the HTTP based repository** and the **config file that is used by the Benchmark** as a source of settings. _Once downloaded the the script runs the benchmark_.
 
     **Configuration file** is generated in the format **compatible with the Benchmark** - for example, if the benchmark is a BASH script, configuration format is BASH array - _various kind of format should be supported, for example .ini file syntax_. Job creator itself has a **facility for translating from XML to settings file** - this means that when creating Benchmarks you don't have to worry about creating a method for translating from XML to configuration file, rather than providing information about the format you want to have it translated.

**After creating the Job**, then new job is then **addedd to the waiting to be executed job queue** of Job Scheduler.

### Job Scheduler
This service monitors a queue and executes new jobs: it reads the Job XML and sets up the Task Scheduler (for example Quartz) to execute them. The reason for leveraging onto a dedicated Task Scheduler is to benefit of advanced features of the Task Scheduler such as automatic dependency solving and so on.

### Archiver
This service **looks for returned execution reports** and **stores them in the NoSQL backend**. Same way as Library, this service uses **additional service hooks**, that are **applied before storing documents into the NoSQL database**. For example, an hook could implement Document Level RBAC by sending the document to **Document Level RBAC Service**, so that it adds the Authorization TAGS to the document before storing them into the NoSQL Datastore. Hooks are also used to send document to the **HP Quality Center (ALM) Archiver service**.

### Fetcher
This service is used to **create custom reports** - it **fetches data from the NoSQL** database and assembles reports matching cutom criteria. Same way as Library, this service uses **additional service hooks**, that are applied before creating the report. For example, an hook could implement Document Level RBAC by sending the **Document to Document Level RBAC Service**, so that it filters documents allowing only the documents the user is actually granted to view.
Fetching relies on Project XML meta tag addedd to reports: you can get everything of a particular project as well as focus on a few TestCase results.

### Additional services (optional)
#### Document Level RBAC
This add-on provides a restriction enforcement by checking/settings a Access Level TAGS on documents.

it servers 2 kind of requests:

* **Authorization check requests**: it verifies **Access Level TAGS** to ensure that requestor is granted to get the requested document. If it does not succeed an empty document is returned to the requestor
* **Authorization sets requests**: it adds **Access Level TAGS** to the supplied document, and returns the enriched version of the document to the caller. **Access Level TAGS** to be added are specified in Project XML.

**Access Level TAGS**: the document is classified with **tags identifiing the role and the single user** - **default action** is to **allow** the listed ones, but it is possible to **specify a negate action** for each single role or user.

The service connects to Authenticator to get the list of roles of the user.

##NOTES
* decide wether or not write an agent
* if using an agent, decide if is the agent that connects to the server (to work nicely with NAT) or the opposite
* if using a push design, beware that it becomes not possible to coordinate infrastructure levels job dependecies