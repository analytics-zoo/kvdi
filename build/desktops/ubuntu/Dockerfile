# 10.239.45.10/arda/kvdi-devbox-ubuntu:0.2
# FROM Dockerfile.desktop
ARG BASE_IMAGE=10.239.45.10/arda/kvdi-ubuntu:alpha   
FROM ${BASE_IMAGE}


# Install vscode, maven, git 
RUN sudo apt-get update && \
    sudo apt-get install -y wget maven git-all && \
    apt install -y software-properties-common apt-transport-https curl && \
    curl -sSL https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add - && \
    sudo add-apt-repository -y "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" && \
    sudo apt update && \
    sudo apt install -y code
# Install Java, Scala
RUN wget http://10.239.45.10:8081/repository/raw/jdk/jdk-8u192-linux-x64.tar.gz && \
    gunzip jdk-8u192-linux-x64.tar.gz && \
    tar -xf jdk-8u192-linux-x64.tar -C /opt && \
    rm jdk-8u192-linux-x64.tar && \
    mv /opt/jdk* /opt/jdk8u192 && \
    ln -s /opt/jdk8u192 /opt/jdk && \
    export http_proxy="http://child-prc.intel.com:913/" && \
    export https_proxy="http://child-prc.intel.com:913/" && \ 
    cd / && wget -c https://downloads.lightbend.com/scala/2.11.8/scala-2.11.8.tgz && \
    (cd / && gunzip < scala-2.11.8.tgz)|(cd /opt && tar -xvf -) && \
    rm /scala-2.11.8.tgz && \
    ln -s /opt/scala-2.11.8 /opt/scala
# Install Idea, pycharm, chrome
RUN export http_proxy="http://child-prc.intel.com:913/" && \
    export https_proxy="http://child-prc.intel.com:913/" && \ 
    cd / && wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    sudo apt-get  install -f -y && \
    sudo apt-get  install -y ./google-chrome-stable_current_amd64.deb && \
    rm google-chrome-stable_current_amd64.deb && \
    cd / && wget https://download.jetbrains.com/idea/ideaIC-2022.2.3.tar.gz && \
    tar -zxvf ideaIC-2022.2.3.tar.gz -C /usr/local/share && \
    mv /usr/local/share/idea* /usr/local/share/idea && \
    rm ideaIC-2022.2.3.tar.gz && \
    sudo echo "alias idea='/usr/local/share/idea/bin/idea.sh'" >> /etc/profile && \
    sudo echo "alias idea='/usr/local/share/idea/bin/idea.sh'" >> /etc/bash.bashrc && \
    cd / && wget https://download.jetbrains.com/python/pycharm-community-2022.2.4.tar.gz && \
    tar -zxvf pycharm-community-2022.2.4.tar.gz -C /usr/local/share && \
    mv /usr/local/share/pycharm* /usr/local/share/pycharm && \
    rm pycharm-community-2022.2.4.tar.gz && \
    sudo echo "alias pycharm='/usr/local/share/pycharm/bin/pycharm.sh'" >> /etc/profile && \
    sudo echo "alias pycharm='/usr/local/share/pycharm/bin/pycharm.sh'" >> /etc/bash.bashrc && \
    sudo echo "alias setp='export http_proxy=http://child-prc.intel.com:913 && export https_proxy=http://child-prc.intel.com:913'" >>  /etc/bash.bashrc && \
    sudo echo "alias setp='export http_proxy=http://child-prc.intel.com:913 && export https_proxy=http://child-prc.intel.com:913'" >>  /etc/profile && \
    sudo echo "unsetp='unset http_proxy && unset https_proxy'" >>  /etc/bash.bashrc && \
    sudo echo "unsetp='unset http_proxy && unset https_proxy'" >>  /etc/profile && \
    /bin/bash -c "source /etc/profile" && \
    /bin/bash -c "source /etc/bash.bashrc"


ENV JAVA_HOME /opt/jdk
ENV SCALA_HOME /opt/scala
ENV PATH $PATH:$JAVA_HOME/bin:$SCALA_HOME/bin