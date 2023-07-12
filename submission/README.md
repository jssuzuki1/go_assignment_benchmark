### What's in this folder?

There are seven relevant other documents in this folder:
- regression.go
- regression.exe
- regression_test.go
- go.mod
- go.sum
- regression.py
- regression.r

regression.go is the main file of interest. It includes all of the data wrangling and regression computation.

regression.exe is the binary executable generated from regression.go. In other words, it runs the program. When you open a terminal and change the directory to the folder this program is located and type 'regression.exe' or, if you are using Powershell, '.\regression.exe', terminal should return 4 coefficients and intercepts along with their computation times. 

regression_test.go includes tests of each of the functions present in regression.go to ensure that they are functional.

regression.py is the python variant of regression.go. It contains the same regression calculations with their total computation time.  

regression.r is the python variant of regression.go. It contains the same regression calculations with their total computation time. 

Daniel said that I didn't need to submit the Python and R programs associated with these, but I had already ammended them to keep track of the amount of time it takes to run those regressions. So I decided to include them anyways to compare performance.

### Perfomance

Cutting right to the chase, Go is crazy fast. Where R and Python take 0.033 and 0.015 seconds to run these four regressions respectively, Go takes a little more than 0.001 seconds to run on my computer with identical results.

### Recommendation to Management

There is a spectre haunting data science-- a spectre of GoLang. Long live the new programming regime, which is harder, better, faster, and stronger. The inefficiency of L'Ancien Régime of R and Python are no more.

To be a bit more serious, I am personally miffed at the fact that the coefficients and slopes have to be manually calculated by the programmer here. Perhaps this is because I am not aware of what other packages may exist. And if people at this company do not have a strong grasp of object-oriented programming, adopting GoLang is going to be quite difficult to say the least.

So my recommendation is the following: if management is willing to set aside some serious time to retrain its staff and transition to this programming language, Go will pay some serious dividends. However, if this short-term cost is too large, remain with the more popular and intuitive languages of Python and Go.