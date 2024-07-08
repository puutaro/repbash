#!/bin/bash


function echoGit(){
echo git
}
#!/bin/bash


function echo3(){
  echo 3
}
#!/bin/bash


function echo1(){
  echo 1
}
#!/bin/bash


function echo2(){
  echo 2
}
#!/bin/bash

### LABELING_SECTION_START
### LABELING_SECTION_END


### SETTING_SECTION_START
terminalDo=ON
openWhere=CW
terminalFocus=OFF
editExecute=ONCE
setVariableTypes="mode:CB=TEST!RUN!BUILD"
beforeCommand=
afterCommand=
execBeforeCtrlCmd=
execAfterCtrlCmd=
appIconPath=
scriptFileName=repBashRunOrTest.sh
### SETTING_SECTION_END


### CMD_VARIABLE_SECTION_START
# mode="BB"
# // mode="aaa"
mode="BUILD"
### CMD_VARIABLE_SECTION_END

### Please write bellow with shell script

working_dir="$HOME/Desktop/share/android/cmds/repbash"

cd "${working_dir}"

export REPLACE_VARIABLES_TSV_RELATIVE_PATH="settingVariables/replaceVariablesTable.tsv"
export APP_ROOT_PATH="/storage/emulated/0/Documents/cmdclick"
export APP_DIR_PATH="/storage/emulated/0/Documents/cmdclick/AppDir"
export UBUNTU_SERVICE_TEMP_DIR_PATH="/storage/emulated/0/Documents/cmdclick/temp/ubuntuService"
export UBUNTU_ENV_TSV_NAME="ubuntu_env.tsv"
export MONITOR_DIR_PATH="/storage/emulated/0/Documents/conf/monitor"
export HTTP2_SHELL_PATH="/storage/emulated/0/Documents/cmdclick/temp/cmd/cmd.sh"
export HTTP2_SHELL_PORT="15000"

echo "echo 'mode BUILD'"
echo "echo 'REPLACE_VARIABLES_TSV_RELATIVE_PATH settingVariables/replaceVariablesTable.tsv'"
echo "echo 'APP_ROOT_PATH /storage/emulated/0/Documents/cmdclick'"
echo "echo 'APP_DIR_PATH /storage/emulated/0/Documents/cmdclick/AppDir'"
echo "echo 'UBUNTU_SERVICE_TEMP_DIR_PATH /storage/emulated/0/Documents/cmdclick/temp/ubuntuService'"
echo "echo 'UBUNTU_ENV_TSV_NAME /suppport/ubuntu_env_temp.tsv'"
echo "echo 'MONITOR_DIR_PATH' /storage/emulated/0/Documents/conf/monitor"
echo "echo 'HTTP2_SHELL_PATH' /storage/emulated/0/Documents/cmdclick/temp/cmd/cmd.sh"
echo "echo 'HTTP2_SHELL_PORT' 15000"
echo "echo 'valName1 vv1'"
echo "echo 'valName2 vv2'"
echo "echo 'valName3 '"
echo "echo 'repbash'"
echo "echo 'cmdMusicPlayerDirPath' /home/dummy/dir/path"
echo "echo 'cmdMusicPlayerListDirPath' /home/dummy/dir/path/list"
echo "echo 'cmdMusicPlayerDirListFilePath' /home/dummy/dir/path/list/music_dir_list"
echo "echo 'cmdYoutubePlayerBlankVal' "
echo "echo 'REPBASH_AGS_CON' valName1=vv1,valName2=vv2,valName3="
echo "'UBUNTU_ENV_TSV_VALUE' ubuntu env tsv value"
### REPBASH_CON