#include "header.h"
namespace fs = std::filesystem;

char *getHostname()
{
    char *computerName = (char *)malloc(1024);
#ifdef _WIN32 || _WIN64
    DWORD size = sizeof(computerName) / sizeof(computerName[0]);
    if (GetComputerNameA(computerName, &size)) {
        return computerName;
    } else {
        return "unknown";
    }
#elif __APPLE__ || __MACH__ || __unix || __unix__ || __linux || __FreeBSD__
    if (gethostname(computerName, sizeof(computerName)) == 0) {
        return computerName;
    }
#else
    return "Unknown";
#endif
return "Unknown";
}

char *getUserName()
{
    char *userName = (char * )malloc(1024);
#ifdef _WIN32 || _WIN64
    DWORD size = sizeof(userName) / sizeof(userName[0]);
    if (GetUserNameA(userName, &size)) {
        return userName;
    } else {
        return "unknown";
    }
#elif __APPLE__ || __MACH__ || __unix || __unix__ || __linux || __FreeBSD__
    if (getlogin_r(userName, sizeof(userName)) == 0)
        return userName;
#else
    return "Unknown";
#endif
return "Unknown";
}


void shiftArr(std::array<float, 10> &values, float newValue)
{
    for (int i = 0; i < 9; i++)
    {
        values[i] = values[i + 1];
    }
    values[9] = newValue;
}

int getNumTasks()
{
    int count = 0;

    for (const auto& entry : fs::directory_iterator("/proc")) {
        if (entry.is_directory()) {
            std::string filename = entry.path().filename().string();
            if (!filename.empty() && std::all_of(filename.begin(), filename.end(), ::isdigit)) {
                std::ifstream status_file(entry.path() / "status");
                if (status_file.is_open()) {
                    std::string line;
                    while (std::getline(status_file, line)) {
                        if (line.find("Name:") == 0) {
                            count++;
                            break;
                        }
                    }
                }
            }
        }
    }
    return count;
}

SystemInfo *getSystemInfo()
{
    SystemInfo *info = (SystemInfo *)malloc(sizeof(SystemInfo));
    info->osName = (char *)malloc(sizeof(getOsName()));
    info->osName =(char *)getOsName();
    info->hostname = getHostname();
    info->numTasks = getNumTasks();
    info->cpu = CPUinfo();
    info->user = getUserName();
    return info;
}