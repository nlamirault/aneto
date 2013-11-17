#
# Copyright 2013 Nicolas Lamirault <nicolas.lamirault@gmail.com>.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.
#

from aneto.config import get_configuration_filename
from cliff.command import Command
import logging
import os


class Backup(Command):
    "A command that backup directories."

    log = logging.getLogger(__name__)

    def take_action(self, parsed_args):
        self.app.stdout.write('Backup data.\n')
        cmd = "rsync -avz --exclude '/AppData/'" + \
              "/cygdrive/c/Users/saltycrane" + \
              "/cygdrive/f/backup/Users"
        os.system(cmd)
