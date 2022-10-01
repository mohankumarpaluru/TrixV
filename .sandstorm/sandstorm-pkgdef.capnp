@0xfd1c6ca5525e2cdb;

using Spk = import "/sandstorm/package.capnp";
# This imports:
#   $SANDSTORM_HOME/latest/usr/include/sandstorm/package.capnp
# Check out that file to see the full, documented package definition format.

const pkgdef :Spk.PackageDefinition = (
  # The package definition. Note that the spk tool looks specifically for the
  # "pkgdef" constant.

  id = "70xjeuj6t85v3s9a3up583q97wu044mnt8aekf3pqs41e3ct2v70",
  # Your app ID is actually its public key. The private key was placed in
  # your keyring. All updates must be signed with the same key.

  manifest = (
    # This manifest is included in your app package to tell Sandstorm
    # about your app.

    appTitle = (defaultText = "Tube"),

    appVersion = 1,  # Increment this for every release.

    appMarketingVersion = (defaultText = "1.1.13"),
    # Human-readable representation of appVersion. Should match the way you
    # identify versions of your app in documentation and marketing.

    actions = [
      # Define your "new document" handlers here.
      ( nounPhrase = (defaultText = "channel"),
        command = .myCommand
        # The command to run when starting for the first time. (".myCommand"
        # is just a constant defined at the bottom of the file.)
      )
    ],

    continueCommand = .myCommand,
    # This is the command called to start your app back up after it has been
    # shut down for inactivity. Here we're using the same command as for
    # starting a new instance, but you could use different commands for each
    # case.

    metadata = (
      icons = (
        appGrid = (svg = embed "../static/tube.svg"),
        grain = (svg = embed "../static/tube.svg"),
        market = (svg = embed "../static/tube.svg"),
        marketBig = (svg = embed "../static/tube.svg"),
      ),

      website = "https://git.mills.io/prologic/tube",

      codeUrl = "https://git.mills.io/prologic/tube",

      license = (openSource = mit),

      categories = [media],

      author = (
        contactEmail = "inbox@jacobweisz.com",

        pgpSignature = embed "pgp-signature",

        upstreamAuthor = "James Mills",
      ),

      pgpKeyring = embed "pgp-keyring",

      description = (defaultText = embed "DESCRIPTION.md"),

      shortDescription = (defaultText = "Simple video sharing"),

      screenshots = [
        (width = 2624, height = 1624, png = embed "../screenshot-1.png"),
        (width = 2624, height = 1624, png = embed "../screenshot-2.png"),
      ],
      changeLog = (defaultText = embed "../CHANGELOG.md"),
    ),
  ),

  sourceMap = (
    # Here we defined where to look for files to copy into your package. The
    # `spk dev` command actually figures out what files your app needs
    # automatically by running it on a FUSE filesystem. So, the mappings
    # here are only to tell it where to find files that the app wants.
    searchPath = [
      ( sourcePath = "." ),  # Search this directory first.
      ( sourcePath = "/",    # Then search the system root directory.
        hidePaths = [ "home", "proc", "sys",
                      "etc/passwd", "etc/hosts", "etc/host.conf",
                      "etc/nsswitch.conf", "etc/resolv.conf" ]
        # You probably don't want the app pulling files from these places,
        # so we hide them. Note that /dev, /var, and /tmp are implicitly
        # hidden because Sandstorm itself provides them.
      )
    ]
  ),

  fileList = "sandstorm-files.list",
  # `spk dev` will write a list of all the files your app uses to this file.
  # You should review it later, before shipping your app.

  alwaysInclude = [],
  # Fill this list with more names of files or directories that should be
  # included in your package, even if not listed in sandstorm-files.list.
  # Use this to force-include stuff that you know you need but which may
  # not have been detected as a dependency during `spk dev`. If you list
  # a directory here, its entire contents will be included recursively.

  bridgeConfig = (
    # Used for integrating permissions and roles into the Sandstorm shell
    # and for sandstorm-http-bridge to pass to your app.
    # Uncomment this block and adjust the permissions and roles to make
    # sense for your app.
    # For more information, see high-level documentation at
    # https://docs.sandstorm.io/en/latest/developing/auth/
    # and advanced details in the "BridgeConfig" section of
    # https://github.com/sandstorm-io/sandstorm/blob/master/src/sandstorm/package.capnp
    viewInfo = (
      # For details on the viewInfo field, consult "ViewInfo" in
      # https://github.com/sandstorm-io/sandstorm/blob/master/src/sandstorm/grain.capnp
  
      permissions = [
      # Permissions which a user may or may not possess.  A user's current
      # permissions are passed to the app as a comma-separated list of `name`
      # fields in the X-Sandstorm-Permissions header with each request.
      #
      # IMPORTANT: only ever append to this list!  Reordering or removing fields
      # will change behavior and permissions for existing grains!  To deprecate a
      # permission, or for more information, see "PermissionDef" in
      # https://github.com/sandstorm-io/sandstorm/blob/master/src/sandstorm/grain.capnp
        (
          name = "admin",
          # Name of the permission, used as an identifier for the permission in cases where string
          # names are preferred.  Used in sandstorm-http-bridge's X-Sandstorm-Permissions HTTP header.
  
          title = (defaultText = "admin"),
          # Display name of the permission, e.g. to display in a checklist of permissions
          # that may be assigned when sharing.
  
          description = (defaultText = "grants ability to do anything"),
          # Prose describing what this role means, suitable for a tool tip or similar help text.
        ),
		(
          name = "upload",
          title = (defaultText = "upload"),
          description = (defaultText = "ability to upload content"),
        ),
      ],
      roles = [
        # Roles are logical collections of permissions.  For instance, your app may have
        # a "viewer" role and an "editor" role
		(
          title = (defaultText = "administrator"),
          # Name of the role.  Shown in the Sandstorm UI to indicate which users have which roles.
  
          permissions  = [true,true],
          # An array indicating which permissions this role carries.
          # It should be the same length as the permissions array in
          # viewInfo, and the order of the lists must match.
  
          verbPhrase = (defaultText = "can do anything"),
          # Brief explanatory text to show in the sharing UI indicating
          # what a user assigned this role will be able to do with the grain.
  
          description = (defaultText = "administrators can edit and configure all settings."),
          # Prose describing what this role means, suitable for a tool tip or similar help text.
        ),
        (
          title = (defaultText = "uploader"),
          permissions  = [false,true],
          verbPhrase = (defaultText = "can upload videos"),
          description = (defaultText = "uploaders may upload new videos."),
        ),
        (
          title = (defaultText = "viewer"),
          permissions  = [false,false],
          verbPhrase = (defaultText = "can watch videos"),
          description = (defaultText = "viewers may watch videos on the channel."),
        ),
      ],
    ),
    #apiPath = "/api",
    # Apps can export an API to the world.  The API is to be used primarily by Javascript
    # code and native apps, so it can't serve out regular HTML to browsers.  If a request
    # comes in to your app's API, sandstorm-http-bridge will prefix the request's path with
    # this string, if specified.
  ),
);

const myCommand :Spk.Manifest.Command = (
  # Here we define the command used to start up your server.
  argv = ["/sandstorm-http-bridge", "8000", "--", "/bin/bash", "/opt/app/.sandstorm/launcher.sh"],
  environ = [
    # Note that this defines the *entire* environment seen by your app.
    (key = "PATH", value = "/usr/local/bin:/usr/bin:/bin"),
    (key = "SANDSTORM", value = "1"),
    # Export SANDSTORM=1 into the environment, so that apps running within Sandstorm
    # can detect if $SANDSTORM="1" at runtime, switching UI and/or backend to use
    # the app's Sandstorm-specific integration code.
  ]
);
