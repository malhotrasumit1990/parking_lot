
1. Code support inputs both from file and Command Line.

2. Folder Structure for parking_lot application : 
		
	parking_lot:
	
		--> src
			--> main.go
			--> MakeFile
			--> glide.yaml
			--> dataLayer
				--> dataAccessLogic.go
			--> businesslayer
				--> businessLogic.go
				--> businessLogic_test.go
				--> constants.go
				--> util.go
				
		--> bin
			-->setup.sh
			-->parking_slot.sh
			-->input.txt


3. How to setup and Run the application:
	
	1. Open terminal and go to "bin" folder
	2. Execute "bash setup.sh"
		
		//This will execute make file, which in turn clears glide cache, downloads dependencies and execute unit test cases. 
		// It also lets you know the code coverage from unit tests.

		
	Steps 1 and 2 are common for both command line execution of application or file based execution of application.

	3.  To test the application from command line : 
		> Go to src folder
		> Execute "go run main.go"	
		> Applcation is now ready to take user inputs from console

	4. To Test the application from File based inputs : 
		> Go to bin folder
		> Execute "bash parking_lot.sh"
		> It will prompt/ask for input file path. Provide the entire path including file name and extension
		> Press Enter

