import { Html } from '@react-email/html';
import { Container } from '@react-email/container';
import { Img } from '@react-email/img';
import { Head } from '@react-email/head';
import { Text } from '@react-email/text';
import { Preview } from '@react-email/preview';
import { Section } from '@react-email/section';

export const CheckoutDone = () => {
  return (
    <Html>
      <Head>
        <title>
          Coffee delivery remix
        </title>
      </Head>
      <Preview>
        Congratulations, thank you for shopping with delivery remix
      </Preview>
      <Section>
        <Container>
          <Section style={{ marginBottom: '24px' }}>
            <Img src="assets/svg/Logo.svg" alt="Logo coffee delivery" width="85" height="40" />
          </Section>

          <Text style={{
            fontFamily: "Baloo 2",
            fontWeight: 800,
            fontSize: '24px',
            lineHeight: '34.4px',
          }}>
            Thank you for shopping with delivery remix
          </Text>
          <Text style={{
            fontFamily: 'Roboto',
            fontWeight: 400,
            fontSize: '16px',
            lineHeight: '20.8px',
          }}>
            Your coffees are being processed and as soon as we confirm the payment they will be shipped and you will be notified. We will also release a code for you to track your purchases.
          </Text>
        </Container>
      </Section>
    </Html>
  );
};