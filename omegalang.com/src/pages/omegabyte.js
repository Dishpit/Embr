import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import OmegaByte2024Info from '@site/src/components/OmegaByte2024';

import Heading from '@theme/Heading';
import styles from './index.module.css';
import OmegaByte2024 from '../../static/img/omegabyte-2024-02.png'

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <img src={OmegaByte2024} />
      </div>
    </header>
  );
}

export default function Home() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title} | OmegaByte`}
      description="The official hackathon for Omega">
      <HomepageHeader />
      <main>
        <OmegaByte2024Info />
      </main>
    </Layout>
  );
}
