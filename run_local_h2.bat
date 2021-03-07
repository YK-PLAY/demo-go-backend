@echo off
SET JAVA_HOME=C:\Users\ztkmk\.jdks\adopt-openjdk-1.8.0_252
SET H2_HOME=C:\Tools\Jar\h2-1.4.200.jar

@REM org.h2.tools.Console
@REM %JAVA_HOME%/bin/java.exe -cp "%H2_HOME%;%H2DRIVERS%;%CLASSPATH%"  org.h2.tools.Server %* -tcp   -tcpAllowOthers -ifNotExists -trace
%JAVA_HOME%/bin/java.exe -cp "%H2_HOME%" org.h2.tools.Server -webAllowOthers -tcpAllowOthers -ifNotExists