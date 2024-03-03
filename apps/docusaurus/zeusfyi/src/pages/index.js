import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';

import styles from './index.module.css';
import HomepageFeatures from "../components/HomepageFeatures";

function HomepageHeader() {
    const {siteConfig} = useDocusaurusContext();
    return (
        <header className={clsx('hero hero--primary', styles.heroBanner)}>
            <div className="container">
                <h1 className="hero__title">{siteConfig.title}</h1>
                <p className="hero__subtitle">{siteConfig.tagline}</p>
                <div className={styles.buttons}>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/mockingbird/intro">
                        Mockingbird
                    </Link>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/zk8s/intro">
                        Developer Platform
                    </Link>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/lb/intro">
                        Adaptive Load Balancer
                    </Link>
                </div>
            </div>
        </header>
    );
}

export default function Home() {
    const {siteConfig} = useDocusaurusContext();
    return (
        <Layout
            wrapperClassName={styles.backgroundHome}
            title={`${siteConfig.title} Documentation`}
            description="Zeusfyi Documentation">
            <HomepageHeader/>
            <main>
                <HomepageFeatures/>
            </main>
        </Layout>
    );
}
