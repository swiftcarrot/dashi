module.exports = {
  title: 'Dashi',
  tagline: 'GraphQL framework',
  url: 'https://dashi.swiftcarrot.dev',
  baseUrl: '/',
  favicon: 'img/favicon.ico',
  organizationName: 'swiftcarrot',
  projectName: 'dashi',
  themeConfig: {
    navbar: {
      title: 'My Site',
      logo: {
        alt: 'My Site Logo',
        src: 'img/logo.svg',
      },
      links: [
        {
          to: 'docs/',
          activeBasePath: 'docs',
          label: 'Docs',
          position: 'left',
        },
        { to: 'blog', label: 'Blog', position: 'left' },
        {
          href: 'https://github.com/swiftcarrot/dashi',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Style Guide',
              to: 'docs/',
            },
            {
              label: 'Second Doc',
              to: 'docs/doc2/',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Stack Overflow',
              href: 'https://stackoverflow.com/questions/tagged/docusaurus',
            },
            {
              label: 'Discord',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/docusaurus',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Blog',
              to: 'blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/swiftcarrot/dashi',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} My Project, Inc. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
          googleAnalytics: {
            trackingID: 'UA-106179075-5',
          },
        },
        docs: {
          homePageId: 'installation',
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://github.com/swiftcarrot/dashi/edit/master/website/',
        },
        blog: {
          showReadingTime: true,
          editUrl: 'https://github.com/swiftcarrot/dashi/edit/master/website/',
        },
      },
    ],
  ],
};
