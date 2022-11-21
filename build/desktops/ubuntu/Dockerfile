# 10.239.45.10/arda/kvdi-devbox-ubuntu:0.1
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
# Install Idea, pycharm
RUN export http_proxy="http://child-prc.intel.com:913/" && \
    export https_proxy="http://child-prc.intel.com:913/" && \ 
    cd / && wget https://download.jetbrains.com/idea/ideaIU-2022.2.3.tar.gz && \
    tar -zxvf ideaIU-2022.2.3.tar.gz -C /usr/local/share && \
    rm ideaIU-2022.2.3.tar.gz && \
    sudo echo "alias idea='/usr/local/share/idea-IU-222.4345.14/bin/idea.sh'" >> /etc/profile && \
    cd / && wget https://download.jetbrains.com/python/pycharm-professional-2022.2.4.tar.gz && \
    tar -zxvf pycharm-professional-2022.2.4.tar.gz -C /usr/local/share && \
    rm pycharm-professional-2022.2.4.tar.gz && \
    sudo echo "alias pycharm='/usr/local/share/pycharm-2022.2.4/bin/pycharm.sh'" >> /etc/profile && \
    /bin/bash -c "source /etc/profile"

ENV JAVA_HOME /opt/jdk
ENV SCALA_HOME /opt/scala
ENV PATH $PATH:$JAVA_HOME/bin:$SCALA_HOME/bin