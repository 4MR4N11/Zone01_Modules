#include "header.h"

const char *getHostname()
{
    std::vector<char> computerName;
#ifdef _WIN32 || _WIN64
    DWORD size = 0;
    GetComputerNameExA(ComputerNameDnsHostname, nullptr, &size);  // Get required size
    if (size > 0) {
        std::vector<char> computerName(size);
        if (GetComputerNameExA(ComputerNameDnsHostname, computerName.data(), &size)) {
            return computerName.data();
        } else {
            return "unknown";
        }
    }
#elif __APPLE__ || __MACH__ || __unix || __unix__ || __linux || __FreeBSD__
    long max_size = sysconf(_SC_HOST_NAME_MAX);
    if (max_size == -1) max_size = 255;
    computerName.resize(max_size + 1);
    
    if (gethostname(computerName.data(), computerName.size()) == 0) {
        return computerName.data();
    }
#else
    return "Unknown";
#endif
return "Unknown";
}

const char *getUserName()
{
    std::vector<char> userName;
#ifdef _WIN32 || _WIN64
    DWORD size = 0;
    GetUserNameA(nullptr, &size);  // Get required size
    if (size > 0) {
        std::vector<char> userName(size);
        if (GetUserNameA(userName.data(), &size)) {
            return userName.data();
        } else {
            return "unknown";
        }
    }
#elif __APPLE__ || __MACH__ || __unix || __unix__ || __linux || __FreeBSD__
    long max_size = sysconf(_SC_LOGIN_NAME_MAX);
    if (max_size == -1) max_size = 255;
    userName.resize(max_size + 1);

    if (getlogin_r(userName.data(), userName.size()) == 0)
        return userName.data();
#else
    return "Unknown";
#endif
return "Unknown";
}

int getNumTasks()
{
    int numTasks = 0;
    std::ifstream file("/proc/stat");
    if (file.is_open())
    {
        std::string line;
        while (std::getline(file, line))
        {
            if (line.find("procs_") == 0)
            {
                int value;
                if (sscanf(line.c_str(), "%*s %d", &value) == 1) {
                    numTasks += value;
                }
            }
        }
        file.close();
    }
    return numTasks;
}

SystemInfo *getSystemInfo()
{
    SystemInfo *info = (SystemInfo *)malloc(sizeof(SystemInfo));
    info->osName = getOsName();
    info->hostname = getHostname();
    info->numTasks = getNumTasks();
    info->cpu = CPUinfo();
    info->user = getUserName();
    return info;
}