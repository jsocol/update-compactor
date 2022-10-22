import compactor
import unittest
from google.protobuf.json_format import MessageToDict
from google.protobuf import field_mask_pb2
import proto


class CompactorTests(unittest.TestCase):
    def test_compactor(self):
        proto_entity = proto.Person(
            id=5432,
            name='Miles',
            address=proto.Address(
                street1='123 Fake St',
                city='Queens',
                postal_code='12345',
                state_code='NY',
            ),
        )

        actual = MessageToDict(proto_entity)

        update1 = proto.UpdatePerson(
            id=5432,
            person=proto.Person(
                address=proto.Address(
                    street2='Apt 78C',
                ),
            ),
            update_mask=field_mask_pb2.FieldMask(
                paths=['address.street2'],
            )
        )

        update2 = proto.UpdatePerson(
            id=5432,
            person=proto.Person(
                emails=['miles.morales@example.com'],
            ),
            update_mask=field_mask_pb2.FieldMask(
                paths=['emails'],
            ),
        )

        update3 = proto.UpdatePerson(
            id=5432,
            person=proto.Person(
                name='Miles Morales',
                address=proto.Address(
                    street1='234 Fake St',
                    state_code='NJ',
                ),
            ),
            update_mask=field_mask_pb2.FieldMask(
                paths=['name', 'address.street1'],
            ),
        )

        expected = {
            'id': 5432,
            'name': 'Miles Morales',
            'address': {
                'street1': '234 Fake St',
                'street2': 'Apt 78C',
                'city': 'Queens',
                'postalCode': '12345',
                'stateCode': 'NY',
            },
            'emails': ['miles.morales@example.com'],
        }

        for update_msg in [update1, update2, update3]:
            updates = compactor.find_updates_from_msg(update_msg)
            compactor.update_entity(actual, updates)

        self.assertDictEqual(expected, actual)

    def test_partial_initial(self):
        proto_entity = proto.Person(
            id=5432,
            name='Miles',
        )

        actual = MessageToDict(proto_entity)

        update1 = proto.UpdatePerson(
            id=5432,
            person=proto.Person(
                address=proto.Address(
                    street1='1234 Fake Blvd',
                    street2='Apt 78C',
                ),
            ),
            update_mask=field_mask_pb2.FieldMask(
                paths=['address.street2', 'address.street1'],
            )
        )

        update2 = proto.UpdatePerson(
            id=5432,
            person=proto.Person(
                emails=['miles.morales@example.com'],
            ),
            update_mask=field_mask_pb2.FieldMask(
                paths=['emails'],
            ),
        )

        expected = {
            'id': 5432,
            'name': 'Miles',
            'address': {
                'street1': '1234 Fake Blvd',
                'street2': 'Apt 78C',
            },
            'emails': ['miles.morales@example.com'],
        }

        for update_msg in [update1, update2]:
            updates = compactor.find_updates_from_msg(update_msg)
            compactor.update_entity(actual, updates)

        self.assertDictEqual(expected, actual)
