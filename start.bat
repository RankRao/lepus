@echo off
setlocal

:: ���ô���ҳΪ936��GBK��
chcp 936 >nul

:: ��ȡ��ǰ�������ļ����ڵ�Ŀ¼
set APP_PATH=%~dp0

:: ȥ��·��ĩβ�ķ�б��
if "%APP_PATH:~-1%" == "\" set APP_PATH=%APP_PATH:~0,-1%

:: ���÷�������
set SERVICE_NAME=LepusService


:: ���÷�������
set SERVICE_DESCRIPTION=LEPUS���ݿ��ع����Զ���ƽ̨

:: ���ÿ�ִ���ļ��Ͳ���
set APP_EXE=%APP_PATH%\lepus.exe
set APP_PARAMS=-c "%APP_PATH%\setting.yml"

:: ���·���Ͳ������е���
echo ��ǰĿ¼: %APP_PATH%
echo ��ִ���ļ�: %APP_EXE%
echo ����: %APP_PARAMS%
echo ��������: %SERVICE_NAME%
echo ��������: %SERVICE_DESCRIPTION%


:: ��������
echo �������� %SERVICE_NAME%...
%APP_EXE% %APP_PARAMS%
if %errorlevel% neq 0 (
    echo ��������ʧ�ܡ��������: %errorlevel%
   
) else (
    echo ���������ɹ���
)


:: ��ͣ�Է�ֹ���ڹر�
echo ��������رմ˴���...
pause >nul