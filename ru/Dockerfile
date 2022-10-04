ARG VERSION_GHOST

FROM ghost:${VERSION_GHOST}-alpine

ARG VERSION_GHOST
ENV VERSION=${VERSION_GHOST}

COPY ./src/content-cta.hbs /var/lib/ghost/versions/$VERSION/core/frontend/helpers/tpl/
COPY ./src/ghost* /var/lib/ghost/versions/$VERSION/core/built/admin/assets/
COPY ./src/vendor* /var/lib/ghost/versions/$VERSION/core/built/admin/assets/
COPY ./src/maintenance.html /var/lib/ghost/versions/$VERSION/core/server/views/
COPY ./src/user.js /var/lib/ghost/versions/$VERSION/core/server/models/
COPY ./src/mail/GhostMailer.js /var/lib/ghost/versions/$VERSION/core/server/services/mail/
COPY ./src/mail/templates/* /var/lib/ghost/versions/$VERSION/core/server/services/mail/templates/
COPY ./src/members/api.js /var/lib/ghost/versions/$VERSION/core/server/services/members/
COPY ./src/members/emails/* /var/lib/ghost/versions/$VERSION/core/server/services/members/emails/
