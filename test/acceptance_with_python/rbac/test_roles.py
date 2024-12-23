import pytest
import weaviate
from _pytest.fixtures import SubRequest
from weaviate.rbac.models import RBAC

from .conftest import _sanitize_role_name, role_wrapper, RoleWrapperProtocol

pytestmark = pytest.mark.xdist_group(name="rbac")


def test_rbac_viewer_assign(
    request: SubRequest, admin_client, viewer_client, role_wrapper: RoleWrapperProtocol
):
    name = _sanitize_role_name(request.node.name)

    admin_client.collections.delete(name)
    admin_client.collections.create(name=name)

    # cannot delete with viewer permissions
    with pytest.raises(weaviate.exceptions.UnexpectedStatusCodeException) as e:
        viewer_client.collections.delete(name)

    # with extra role that has those permissions it works
    with role_wrapper(
        admin_client,
        request,
        RBAC.permissions.collections(collection=name, delete_collection=True),
        "viewer-user",
    ):
        viewer_client.collections.delete(name)

    admin_client.collections.delete(name)


def test_rbac_with_regexp(
    request: SubRequest, admin_client, custom_client, role_wrapper: RoleWrapperProtocol
):
    name = _sanitize_role_name(request.node.name)
    base = "python_"
    python_name = base + name
    admin_client.collections.delete([name, python_name])
    admin_client.collections.create(name=name)
    admin_client.collections.create(name=python_name)

    with pytest.raises(weaviate.exceptions.UnexpectedStatusCodeException) as e:
        custom_client.collections.delete(python_name)

    # can delete everything starting with "python_" but nothing else
    required_permissions = [
        RBAC.permissions.collections(collection="*", read_config=True),
        RBAC.permissions.collections(collection=base + "*", delete_collection=True),
    ]
    with role_wrapper(admin_client, request, required_permissions):
        custom_client.collections.delete(python_name)
        with pytest.raises(weaviate.exceptions.UnexpectedStatusCodeException) as e:
            custom_client.collections.delete(name)
    admin_client.collections.delete([name, python_name])
