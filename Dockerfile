FROM golang:1.19-bullseye

ARG USERNAME=butler
ARG USER_UID=1000
ARG USER_GID=$USER_UID
# Install developemen tools

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

RUN apt install git -y
RUN apt install fish -y

USER $USERNAME

RUN go install github.com/cosmtrek/air@latest

RUN mkdir -p ~/.config/fish/
RUN touch ~/.config/fish/config.fish
RUN echo 'set fish_greeting' >> ~/.config/fish/config.fish

RUN mkdir -p /home/$USERNAME/.vscode-server/extensions \
    && chown -R $USERNAME /home/$USERNAME/.vscode-server

ENV SHELL /usr/bin/fish
