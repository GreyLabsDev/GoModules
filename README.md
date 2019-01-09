# GoModules
Useful modules for my current and future Go projects, it can be used for some usual tasks in different type of projects

## Modules:
- **filetool.go**

- **taskmanager.go**

**Filetool functions:**
- OpenFile(fileName string) *os.File

- CreateFile(fileName string) *os.File

- CloseFile(file *os.File)

- AppendToFile(file *os.File, stringToAppend string)

- AppendToFileAndClose(file *os.File, stringToAppend string)

- ReadImage(imagePath string) (outImage image.Image) - for .png files

- SaveImage(imagePath string, targetImage image.Image) - for .png files

- DownloadImage(imageURL string, imageName string, callback DownloadCallback) - callback executing after image dowloaded

**Taskmanager functions:**
- **StartPeriodicTask(taskExecutionPeriod int64, timePeriodType string, workingPeriodStartHour int, workingPeriodEndHour int, controlChannel chan string, task PeriodicTask)** - executing PeriodicTask (just function) with some parameters periodically, task can be stopped by sending to its controlChannel "stop" string

- **DoPeriodicTaskAtTime(timeToStart string, controlChannel chan string, task PeriodicTask)** - executing same PeriodicTask at current time every day, every minute time checking, time format example "20:30" (HH:mm), task can be stopped by sending to its controlChannel "stop" string

- **CompleteTaskQueue(taskQueue []SingleTask, endMessage string, controlChannel chan string)** - executing task array/slice alternately, one by one, after execunig fuction sends endMessage to it controlChannel, SingleTask - just function
