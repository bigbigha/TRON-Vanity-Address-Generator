@echo off
cls
echo TRON Vanity Address Generator
echo ============================
echo.

:input_loop
echo Please enter the prefix you want to match (or press Enter to skip):
set /p prefix=
echo.

echo Please enter the suffix you want to match (or press Enter to skip):
set /p suffix=
echo.

echo Please enter the number of addresses to generate (default: 1):
set /p count=
if "%count%"=="" set count=1
echo.

echo Searching for TRON vanity addresses...
echo.

:: Build the command
set cmd=tronvanity-windows-amd64.exe
if not "%prefix%"=="" set cmd=%cmd% --prefix %prefix%
if not "%suffix%"=="" set cmd=%cmd% --suffix %suffix%
set cmd=%cmd% --count %count%

:: Execute the command
%cmd%

echo.
echo.
echo Press any key to generate more addresses, or Ctrl+C to exit...
pause >nul
goto input_loop