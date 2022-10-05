#include <windows.h>
#include <string>

INT WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, PSTR lpCmdLine,
                   INT nCmdShow) {

    // Get binary and TMPPATH directories
    std::wstring path(MAX_PATH, 0);
    GetModuleFileNameW(NULL, &path[0], MAX_PATH);
    path.erase(path.rfind('\\')+1);

    std::wstring tmpPath(MAX_PATH, 0);
    GetTempPathW(MAX_PATH, &tmpPath[0]);

    STARTUPINFO si;
    PROCESS_INFORMATION pi;

    ZeroMemory( &si, sizeof(si) );
    si.cb = sizeof(si);
    ZeroMemory( &pi, sizeof(pi) );

    // Command to run
    std::wstring cmdline = path
        + L"DOSbox.exe " \
        + path + L"PRINCE\\PRINCE.EXE " \
        + L"-exit " \
        + L"-noconsole " \
        + L"-fullscreen " 
        + L"-conf " + path + L"DOSbox.conf";
  
    if(!CreateProcessW(
        NULL,
        &cmdline[0],
        NULL,
        NULL,
        FALSE,          // Set handle inheritance to FALSE
        0,
        NULL,           // Use parent's environment block
        &tmpPath[0],    // Start in TMPPATH directory
        &si,
        &pi)
    ) {
        return 1;
    }

    // Wait until child process exits and close handles
    WaitForSingleObject(pi.hProcess, INFINITE);
    CloseHandle(pi.hProcess);
    CloseHandle(pi.hThread);
    return 0;
}