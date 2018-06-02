# What is it?
This is example of the deployable solution that gathers currency data from https://fixer.io/ each day and save it in a custom Golang DB. You can then request convertion of EURO to any other currency through a webhooked bot in slacks. The bot also uses google diaglogflow to handle chatting with users. This code was also made to run within a docker container. 

# How to build
I used jetbrain's Goland to build this project but should build like any other golang application on ubuntu. I didn't make any documentation for this project and it's over 8 months since i actually made this so i don't recall the whole process sadly.

# Who to contact
akselhj@stud.ntnu.no

## Example chat from a deployed version on slacks:
<img src="https://github.com/Avokadoen/CurrencyCloudSolution/blob/master/example.PNG" height="900" width="450" />
