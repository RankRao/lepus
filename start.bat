@echo off
setlocal

:: 设置代码页为936（GBK）
chcp 936 >nul

:: 获取当前批处理文件所在的目录
set APP_PATH=%~dp0

:: 去掉路径末尾的反斜杠
if "%APP_PATH:~-1%" == "\" set APP_PATH=%APP_PATH:~0,-1%

:: 设置服务名称
set SERVICE_NAME=LepusService


:: 设置服务描述
set SERVICE_DESCRIPTION=LEPUS数据库监控管理自动化平台

:: 设置可执行文件和参数
set APP_EXE=%APP_PATH%\lepus.exe
set APP_PARAMS=-c "%APP_PATH%\setting.yml"

:: 输出路径和参数进行调试
echo 当前目录: %APP_PATH%
echo 可执行文件: %APP_EXE%
echo 参数: %APP_PARAMS%
echo 服务名称: %SERVICE_NAME%
echo 服务描述: %SERVICE_DESCRIPTION%


:: 启动服务
echo 启动服务 %SERVICE_NAME%...
%APP_EXE% %APP_PARAMS%
if %errorlevel% neq 0 (
    echo 启动服务失败。错误代码: %errorlevel%
   
) else (
    echo 服务启动成功。
)


:: 暂停以防止窗口关闭
echo 按任意键关闭此窗口...
pause >nul