import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Simple Syntax',
    Svg: require('@site/static/img/undraw_docusaurus_mountain.svg').default,
    description: (
      <>
        Embr&apos;s syntax is straightforward and easy to learn, using clear type annotations and intuitive return type markers, making programming accessible for beginners and efficient for experienced developers.
      </>
    ),
  },
  {
    title: 'Future-Proof Design',
    Svg: require('@site/static/img/undraw_docusaurus_tree.svg').default,
    description: (
      <>
        Embr is built with a future-proof design, ensuring that every specification iteration remains compatible with new releases. Unlike other languages such as Rust, Embr avoids breaking changes and forwards incompatibility, providing a stable and reliable coding experience over time.
      </>
    ),
  },
  {
    title: 'Streamlined Learning Curve',
    Svg: require('@site/static/img/undraw_docusaurus_react.svg').default,
    description: (
      <>
        Embr offers a minimalist feature set, with explicit types and familiar programming concepts, designed to flatten the learning curve and enable quick mastery for new programmers.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
