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
export APP_DIR_PATH="${APP_ROOT_PATH}/AppDir"
export UBUNTU_SERVICE_TEMP_DIR_PATH="${APP_ROOT_PATH}/temp/ubuntuService"
export UBUNTU_ENV_TSV_NAME="/suppport/ubuntu_env_temp.tsv"
export MONITOR_DIR_PATH="/storage/emulated/0/Documents/conf/monitor"
export HTTP2_SHELL_PATH="/storage/emulated/0/Documents/cmdclick/temp/cmd/cmd.sh"
export HTTP2_SHELL_PORT="15000"

echo "echo 'mode ${mode}'"
echo "echo 'REPLACE_VARIABLES_TSV_RELATIVE_PATH ${REPLACE_VARIABLES_TSV_RELATIVE_PATH}'"
echo "echo 'APP_ROOT_PATH ${APP_ROOT_PATH}'"
echo "echo 'APP_DIR_PATH ${APP_DIR_PATH}'"
echo "echo 'UBUNTU_SERVICE_TEMP_DIR_PATH ${UBUNTU_SERVICE_TEMP_DIR_PATH}'"
echo "echo 'UBUNTU_ENV_TSV_NAME ${UBUNTU_ENV_TSV_NAME}'"
echo "echo 'MONITOR_DIR_PATH' ${MONITOR_DIR_PATH}"
echo "echo 'HTTP2_SHELL_PATH' ${HTTP2_SHELL_PATH}"
echo "echo 'HTTP2_SHELL_PORT' ${HTTP2_SHELL_PORT}"
echo "echo 'valName1 ${valName1}'"
echo "echo 'valName2 ${valName2}'"
echo "echo 'repbash'"
echo "echo 'cmdMusicPlayerDirPath' ${cmdMusicPlayerDirPath}"
echo "echo 'cmdMusicPlayerListDirPath' ${cmdMusicPlayerListDirPath}"
echo "echo 'cmdMusicPlayerDirListFilePath' ${cmdMusicPlayerDirListFilePath}"
echo "echo 'cmdYoutubePlayerBlankVal' ${cmdYoutubePlayerBlankVal}"
echo "echo 'REPBASH_AGS_CON' ${REPBASH_ARGS_CON}"

go build -o repbash cmd/repbash/main.go \
;exec ./repbash \
    "${0}" \
    -t \
      '${REPLACE_VARIABLE_TABLE_TSV_PATH}' \
    -a "valName1=vv1,valName2=vv2" \
    -i \
      "\${IMPORT_PATH1}" \
      "\${IMPORT_PATH2}" \
      "\${IMPORT_PATH3}" \
      "\${IMPORT_GIT1}"
